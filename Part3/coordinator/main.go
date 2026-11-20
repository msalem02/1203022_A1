package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	pb "example.com/distributedpow/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Job struct {
	ID          string    `json:"job_id"`
	Data        string    `json:"data"`
	Difficulty  uint32    `json:"difficulty"`
	Status      string    `json:"status"`
	Checked     uint64    `json:"nonces_checked"`
	Elapsed     float64   `json:"elapsed_seconds"`
	Nonce       *uint64   `json:"nonce,omitempty"`
	Hash        string    `json:"hash,omitempty"`
	Started     time.Time `json:"started_at"`
	mu          sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	subscribers map[chan string]struct{}
}

func (j *Job) snapshot() map[string]any {
	j.mu.Lock()
	defer j.mu.Unlock()
	elapsed := time.Since(j.Started).Seconds()
	return map[string]any{"job_id": j.ID, "data": j.Data, "difficulty": j.Difficulty, "status": j.Status, "nonces_checked": j.Checked, "elapsed_seconds": elapsed, "nonce": j.Nonce, "hash": j.Hash}
}
func (j *Job) notify(event string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	for ch := range j.subscribers {
		select {
		case ch <- event:
		default:
		}
	}
}
func (j *Job) event(eventType string) {
	data, _ := json.Marshal(j.snapshot())
	j.notify("event: " + eventType + "\ndata: " + string(data) + "\n\n")
}
func (j *Job) addCount(n uint64) {
	j.mu.Lock()
	if j.Status == "RUNNING" {
		j.Checked += n
	}
	j.mu.Unlock()
}
func (j *Job) finish(nonce uint64, hash string) bool {
	j.mu.Lock()
	if j.Status != "RUNNING" {
		j.mu.Unlock()
		return false
	}
	j.Status = "FOUND"
	j.Nonce = &nonce
	j.Hash = hash
	j.mu.Unlock()
	j.event("final")
	j.cancel()
	return true
}
func (j *Job) stop() bool {
	j.mu.Lock()
	if j.Status != "RUNNING" {
		j.mu.Unlock()
		return false
	}
	j.Status = "CANCELLED"
	j.mu.Unlock()
	j.event("final")
	j.cancel()
	return true
}

type workerInfo struct {
	ID         string    `json:"id"`
	Address    string    `json:"address"`
	Alive      bool      `json:"alive"`
	Assignment string    `json:"assignment"`
	Seen       time.Time `json:"last_heartbeat"`
	Busy       bool      `json:"-"`
}
type coordinator struct {
	pb.UnimplementedCoordinatorServiceServer
	mu               sync.Mutex
	jobs             map[string]*Job
	workers          map[string]*workerInfo
	chunk            uint64
	progressInterval time.Duration
}

func (c *coordinator) RegisterWorker(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if r.WorkerId == "" || r.Address == "" {
		return &pb.RegisterReply{Accepted: false}, nil
	}
	c.mu.Lock()
	w := c.workers[r.WorkerId]
	if w == nil {
		w = &workerInfo{ID: r.WorkerId}
		c.workers[r.WorkerId] = w
	}
	w.Address = r.Address
	w.Alive = true
	w.Seen = time.Now()
	c.mu.Unlock()
	log.Printf("Registered worker %s at %s", r.WorkerId, r.Address)
	return &pb.RegisterReply{Accepted: true}, nil
}
func (c *coordinator) Heartbeat(ctx context.Context, r *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	w := c.workers[r.WorkerId]
	if w == nil {
		return &pb.HeartbeatReply{Ok: false}, nil
	}
	w.Seen = time.Now()
	w.Alive = true
	return &pb.HeartbeatReply{Ok: true}, nil
}
func (c *coordinator) monitor() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for range t.C {
		c.mu.Lock()
		for _, w := range c.workers {
			if time.Since(w.Seen) > 6*time.Second && w.Alive {
				w.Alive = false
				log.Printf("Worker %s heartbeat timeout", w.ID)
			}
		}
		c.mu.Unlock()
	}
}
func (c *coordinator) reserve(jobID string) (string, string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ids := make([]string, 0, len(c.workers))
	for id := range c.workers {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		w := c.workers[id]
		if w.Alive && !w.Busy && time.Since(w.Seen) <= 6*time.Second {
			w.Busy = true
			w.Assignment = jobID
			return id, w.Address, true
		}
	}
	return "", "", false
}
func (c *coordinator) release(id, job string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	w := c.workers[id]
	if w != nil && w.Assignment == job {
		w.Busy = false
		w.Assignment = ""
	}
}

type workRange struct{ start, end uint64 }
type workResult struct {
	r      workRange
	failed bool
	worker string
}

func (c *coordinator) runJob(j *Job) {
	results := make(chan workResult, 1024)
	var next uint64
	retry := make([]workRange, 0)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	progress := time.NewTicker(c.progressInterval)
	defer progress.Stop()
	for {
		if j.ctx.Err() != nil {
			log.Printf("Job %s stopped", j.ID)
			return
		}
		for {
			id, address, ok := c.reserve(j.ID)
			if !ok {
				break
			}
			var r workRange
			if len(retry) > 0 {
				r = retry[0]
				retry = retry[1:]
			} else {
				r = workRange{start: next, end: next + c.chunk}
				if r.end < next {
					r.end = math.MaxUint64
				}
				next = r.end
			}
			go c.search(j, id, address, r, results)
		}
		select {
		case result := <-results:
			c.release(result.worker, j.ID)
			if result.failed && j.ctx.Err() == nil {
				c.mu.Lock()
				if w := c.workers[result.worker]; w != nil {
					w.Alive = false
				}
				c.mu.Unlock()
				retry = append(retry, result.r)
				log.Printf("Job %s reassigning failed range [%d, %d)", j.ID, result.r.start, result.r.end)
			}
		case <-progress.C:
			j.event("progress")
		case <-ticker.C:
		case <-j.ctx.Done():
			return
		}
	}
}
func (c *coordinator) search(j *Job, id, address string, r workRange, results chan<- workResult) {
	failed := true
	defer func() {
		c.release(id, j.ID)
		select {
		case results <- workResult{r: r, failed: failed, worker: id}:
		default:
		}
	}()
	log.Printf("Job %s assigned [%d, %d) to worker %s", j.ID, r.start, r.end, id)
	dialCtx, cancel := context.WithTimeout(j.ctx, 30*time.Second)
	defer cancel()
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Worker %s connection error: %v", id, err)
		return
	}
	defer conn.Close()
	stream, err := pb.NewWorkerServiceClient(conn).SearchRange(dialCtx, &pb.SearchRequest{JobId: j.ID, Data: j.Data, Difficulty: j.Difficulty, Start: r.start, End: r.end})
	if err != nil {
		log.Printf("Worker %s SearchRange error: %v", id, err)
		return
	}
	for {
		update, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("Worker %s stream ended without a result", id)
			} else if j.ctx.Err() == nil {
				log.Printf("Worker %s failed during job %s: %v", id, j.ID, err)
			}
			return
		}
		j.addCount(update.Checked)
		if update.Found {
			h := sha256.Sum256([]byte(j.Data + strconv.FormatUint(update.Nonce, 10)))
			hash := hex.EncodeToString(h[:])
			if hash == update.Hash && strings.HasPrefix(hash, strings.Repeat("0", int(j.Difficulty))) {
				failed = false
				if j.finish(update.Nonce, hash) {
					log.Printf("JOB %s FOUND nonce %d hash %s by %s", j.ID, update.Nonce, hash, id)
				}
				return
			}
			log.Printf("Invalid answer from worker %s", id)
			return
		}
		if update.Finished {
			failed = false
			return
		}
	}
}
func jsonWrite(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func (c *coordinator) handler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	if path == "jobs" {
		switch r.Method {
		case http.MethodPost:
			var body struct {
				Data       *string `json:"data"`
				Difficulty *int    `json:"difficulty"`
			}
			r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			if err := dec.Decode(&body); err != nil || body.Data == nil || *body.Data == "" || body.Difficulty == nil || *body.Difficulty < 1 || *body.Difficulty > 8 {
				jsonWrite(w, 400, map[string]string{"error": "data is required and difficulty must be between 1 and 8"})
				return
			}
			var extra any
			if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
				jsonWrite(w, 400, map[string]string{"error": "invalid JSON body"})
				return
			}
			raw := make([]byte, 8)
			if _, err := rand.Read(raw); err != nil {
				jsonWrite(w, 500, map[string]string{"error": "could not create ID"})
				return
			}
			ctx, cancel := context.WithCancel(context.Background())
			j := &Job{ID: hex.EncodeToString(raw), Data: *body.Data, Difficulty: uint32(*body.Difficulty), Status: "RUNNING", Started: time.Now(), ctx: ctx, cancel: cancel, subscribers: make(map[chan string]struct{})}
			c.mu.Lock()
			c.jobs[j.ID] = j
			c.mu.Unlock()
			log.Printf("New job %s with difficulty %d", j.ID, j.Difficulty)
			go c.runJob(j)
			w.Header().Set("Location", "/jobs/"+j.ID)
			jsonWrite(w, http.StatusAccepted, map[string]string{"job_id": j.ID, "status": "RUNNING"})
		case http.MethodGet:
			c.mu.Lock()
			list := make([]*Job, 0, len(c.jobs))
			for _, j := range c.jobs {
				list = append(list, j)
			}
			c.mu.Unlock()
			out := make([]any, 0, len(list))
			for _, j := range list {
				out = append(out, j.snapshot())
			}
			jsonWrite(w, 200, out)
		default:
			w.Header().Set("Allow", "GET, POST")
			jsonWrite(w, 405, map[string]string{"error": "method not allowed"})
		}
		return
	}
	if path == "workers" {
		if r.Method != http.MethodGet {
			jsonWrite(w, 405, map[string]string{"error": "method not allowed"})
			return
		}
		c.mu.Lock()
		out := make([]workerInfo, 0, len(c.workers))
		for _, worker := range c.workers {
			out = append(out, *worker)
		}
		c.mu.Unlock()
		sort.Slice(out, func(i, k int) bool { return out[i].ID < out[k].ID })
		jsonWrite(w, 200, out)
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) >= 2 && parts[0] == "jobs" && parts[1] != "" {
		c.mu.Lock()
		j := c.jobs[parts[1]]
		c.mu.Unlock()
		if j == nil {
			jsonWrite(w, 404, map[string]string{"error": "job not found"})
			return
		}
		if len(parts) == 3 && parts[2] == "events" && r.Method == http.MethodGet {
			events(w, r, j)
			return
		}
		if len(parts) != 2 {
			jsonWrite(w, 404, map[string]string{"error": "not found"})
			return
		}
		switch r.Method {
		case http.MethodGet:
			jsonWrite(w, 200, j.snapshot())
		case http.MethodDelete:
			if !j.stop() {
				jsonWrite(w, 409, map[string]string{"error": "job already finished"})
				return
			}
			log.Printf("Job %s CANCELLED", j.ID)
			jsonWrite(w, 200, j.snapshot())
		default:
			jsonWrite(w, 405, map[string]string{"error": "method not allowed"})
		}
		return
	}
	jsonWrite(w, 404, map[string]string{"error": "not found"})
}
func events(w http.ResponseWriter, r *http.Request, j *Job) {
	f, ok := w.(http.Flusher)
	if !ok {
		jsonWrite(w, 500, map[string]string{"error": "streaming not supported"})
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	ch := make(chan string, 32)
	j.mu.Lock()
	j.subscribers[ch] = struct{}{}
	finished := j.Status != "RUNNING"
	j.mu.Unlock()
	defer func() { j.mu.Lock(); delete(j.subscribers, ch); j.mu.Unlock() }()
	b, _ := json.Marshal(j.snapshot())
	fmt.Fprintf(w, "event: progress\ndata: %s\n\n", b)
	f.Flush()
	if finished {
		fmt.Fprintf(w, "event: final\ndata: %s\n\n", b)
		f.Flush()
		return
	}
	ping := time.NewTicker(10 * time.Second)
	defer ping.Stop()
	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-ch:
			if _, err := fmt.Fprint(w, msg); err != nil {
				return
			}
			f.Flush()
			if strings.HasPrefix(msg, "event: final") {
				return
			}
		case <-ping.C:
			fmt.Fprint(w, ": ping\n\n")
			f.Flush()
		}
	}
}
func main() {
	httpPort := flag.Int("http-port", 8080, "REST API port")
	grpcPort := flag.Int("grpc-port", 50051, "coordinator gRPC port")
	chunk := flag.Uint64("chunk", 100000, "nonces per worker range")
	interval := flag.Duration("progress", 1*time.Second, "SSE progress interval")
	flag.Parse()
	if *chunk == 0 {
		log.Fatal("chunk must be positive")
	}
	c := &coordinator{jobs: make(map[string]*Job), workers: make(map[string]*workerInfo), chunk: *chunk, progressInterval: *interval}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	gs := grpc.NewServer()
	pb.RegisterCoordinatorServiceServer(gs, c)
	go func() {
		log.Printf("Coordinator gRPC on :%d", *grpcPort)
		if err := gs.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	go c.monitor()
	log.Printf("REST API on http://localhost:%d", *httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), http.HandlerFunc(c.handler)))
}
