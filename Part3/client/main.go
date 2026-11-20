package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	address := flag.String("url", "http://localhost:8080", "REST coordinator URL")
	data := flag.String("data", "hello", "data to hash")
	difficulty := flag.Int("difficulty", 4, "leading zero hex digits")
	flag.Parse()
	payload, _ := json.Marshal(map[string]any{"data": *data, "difficulty": *difficulty})
	response, err := http.Post(*address+"/jobs", "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode != 202 {
		log.Fatalf("Unexpected response: %s %s", response.Status, body)
	}
	var result struct {
		JobID string `json:"job_id"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Submitted job:", result.JobID)
	for {
		time.Sleep(time.Second)
		res, err := http.Get(*address + "/jobs/" + result.JobID)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		var state struct {
			Status  string  `json:"status"`
			Checked uint64  `json:"nonces_checked"`
			Nonce   *uint64 `json:"nonce"`
			Hash    string  `json:"hash"`
		}
		if err := json.Unmarshal(body, &state); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Status: %s  Checked: %d\n", state.Status, state.Checked)
		if state.Status != "RUNNING" {
			fmt.Println("Final response:", string(body))
			return
		}
	}
}
