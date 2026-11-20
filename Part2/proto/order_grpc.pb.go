
package orderpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	OrderService_GetOrder_FullMethodName   = "/order.OrderService/GetOrder"
	OrderService_TrackOrder_FullMethodName = "/order.OrderService/TrackOrder"
)

type OrderServiceClient interface {
	GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderReply, error)
	TrackOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderUpdate], error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderReply)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TrackOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderUpdate], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_TrackOrder_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OrderRequest, OrderUpdate]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_TrackOrderClient = grpc.ServerStreamingClient[OrderUpdate]

type OrderServiceServer interface {
	GetOrder(context.Context, *OrderRequest) (*OrderReply, error)
	TrackOrder(*OrderRequest, grpc.ServerStreamingServer[OrderUpdate]) error
	mustEmbedUnimplementedOrderServiceServer()
}

type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderRequest) (*OrderReply, error) {
	return nil, status.Error(codes.Unimplemented, "method GetOrder not implemented")
}

func (UnimplementedOrderServiceServer) TrackOrder(*OrderRequest, grpc.ServerStreamingServer[OrderUpdate]) error {
	return status.Error(codes.Unimplemented, "method TrackOrder not implemented")
}

func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

func (UnimplementedOrderServiceServer) testEmbeddedByValue() {}

type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TrackOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).TrackOrder(m, &grpc.GenericServerStream[OrderRequest, OrderUpdate]{ServerStream: stream})
}

type OrderService_TrackOrderServer = grpc.ServerStreamingServer[OrderUpdate]

var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TrackOrder",
			Handler:       _OrderService_TrackOrder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/order.proto",
}
