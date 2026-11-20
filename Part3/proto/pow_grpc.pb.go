package powpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)
const _ = grpc.SupportPackageIsVersion9
const (
	CoordinatorService_RegisterWorker_FullMethodName = "/pow.CoordinatorService/RegisterWorker"
	CoordinatorService_Heartbeat_FullMethodName      = "/pow.CoordinatorService/Heartbeat"
)

type CoordinatorServiceClient interface {
	RegisterWorker(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error)
}

type coordinatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorServiceClient(cc grpc.ClientConnInterface) CoordinatorServiceClient {
	return &coordinatorServiceClient{cc}
}

func (c *coordinatorServiceClient) RegisterWorker(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, CoordinatorService_RegisterWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatReply)
	err := c.cc.Invoke(ctx, CoordinatorService_Heartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CoordinatorServiceServer interface {
	RegisterWorker(context.Context, *RegisterRequest) (*RegisterReply, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error)
	mustEmbedUnimplementedCoordinatorServiceServer()
}
type UnimplementedCoordinatorServiceServer struct{}
func (UnimplementedCoordinatorServiceServer) RegisterWorker(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Error(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (UnimplementedCoordinatorServiceServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error) {
	return nil, status.Error(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedCoordinatorServiceServer) mustEmbedUnimplementedCoordinatorServiceServer() {}
func (UnimplementedCoordinatorServiceServer) testEmbeddedByValue()                            {}
type UnsafeCoordinatorServiceServer interface {
	mustEmbedUnimplementedCoordinatorServiceServer()
}

func RegisterCoordinatorServiceServer(s grpc.ServiceRegistrar, srv CoordinatorServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoordinatorService_ServiceDesc, srv)
}

func _CoordinatorService_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServiceServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinatorService_RegisterWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServiceServer).RegisterWorker(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinatorService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinatorService_Heartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var CoordinatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pow.CoordinatorService",
	HandlerType: (*CoordinatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _CoordinatorService_RegisterWorker_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _CoordinatorService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pow.proto",
}

const (
	WorkerService_SearchRange_FullMethodName = "/pow.WorkerService/SearchRange"
)

type WorkerServiceClient interface {
	SearchRange(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchUpdate], error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) SearchRange(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchUpdate], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WorkerService_ServiceDesc.Streams[0], WorkerService_SearchRange_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SearchRequest, SearchUpdate]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}
type WorkerService_SearchRangeClient = grpc.ServerStreamingClient[SearchUpdate]
type WorkerServiceServer interface {
	SearchRange(*SearchRequest, grpc.ServerStreamingServer[SearchUpdate]) error
	mustEmbedUnimplementedWorkerServiceServer()
}
type UnimplementedWorkerServiceServer struct{}
func (UnimplementedWorkerServiceServer) SearchRange(*SearchRequest, grpc.ServerStreamingServer[SearchUpdate]) error {
	return status.Error(codes.Unimplemented, "method SearchRange not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}
func (UnimplementedWorkerServiceServer) testEmbeddedByValue()                       {}
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}
func _WorkerService_SearchRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerServiceServer).SearchRange(m, &grpc.GenericServerStream[SearchRequest, SearchUpdate]{ServerStream: stream})
}
type WorkerService_SearchRangeServer = grpc.ServerStreamingServer[SearchUpdate]
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pow.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchRange",
			Handler:       _WorkerService_SearchRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/pow.proto",
}
