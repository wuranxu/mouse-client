// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: mouse.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MouseServiceClient is the client API for MouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MouseServiceClient interface {
	// connect
	Connect(ctx context.Context, opts ...grpc.CallOption) (MouseService_ConnectClient, error)
	// disconnect
	Disconnect(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MouseResponse, error)
	// stat
	Stat(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MouseResponse, error)
	// start job for scene
	Start(ctx context.Context, in *Task, opts ...grpc.CallOption) (*MouseResponse, error)
	// stop job for scene
	Stop(ctx context.Context, in *StopTask, opts ...grpc.CallOption) (*MouseResponse, error)
}

type mouseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMouseServiceClient(cc grpc.ClientConnInterface) MouseServiceClient {
	return &mouseServiceClient{cc}
}

func (c *mouseServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (MouseService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &MouseService_ServiceDesc.Streams[0], "/MouseService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &mouseServiceConnectClient{stream}
	return x, nil
}

type MouseService_ConnectClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type mouseServiceConnectClient struct {
	grpc.ClientStream
}

func (x *mouseServiceConnectClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mouseServiceConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mouseServiceClient) Disconnect(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MouseResponse, error) {
	out := new(MouseResponse)
	err := c.cc.Invoke(ctx, "/MouseService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mouseServiceClient) Stat(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MouseResponse, error) {
	out := new(MouseResponse)
	err := c.cc.Invoke(ctx, "/MouseService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mouseServiceClient) Start(ctx context.Context, in *Task, opts ...grpc.CallOption) (*MouseResponse, error) {
	out := new(MouseResponse)
	err := c.cc.Invoke(ctx, "/MouseService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mouseServiceClient) Stop(ctx context.Context, in *StopTask, opts ...grpc.CallOption) (*MouseResponse, error) {
	out := new(MouseResponse)
	err := c.cc.Invoke(ctx, "/MouseService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MouseServiceServer is the server API for MouseService service.
// All implementations must embed UnimplementedMouseServiceServer
// for forward compatibility
type MouseServiceServer interface {
	// connect
	Connect(MouseService_ConnectServer) error
	// disconnect
	Disconnect(context.Context, *Message) (*MouseResponse, error)
	// stat
	Stat(context.Context, *Message) (*MouseResponse, error)
	// start job for scene
	Start(context.Context, *Task) (*MouseResponse, error)
	// stop job for scene
	Stop(context.Context, *StopTask) (*MouseResponse, error)
	mustEmbedUnimplementedMouseServiceServer()
}

// UnimplementedMouseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMouseServiceServer struct {
}

func (UnimplementedMouseServiceServer) Connect(MouseService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMouseServiceServer) Disconnect(context.Context, *Message) (*MouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedMouseServiceServer) Stat(context.Context, *Message) (*MouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedMouseServiceServer) Start(context.Context, *Task) (*MouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedMouseServiceServer) Stop(context.Context, *StopTask) (*MouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedMouseServiceServer) mustEmbedUnimplementedMouseServiceServer() {}

// UnsafeMouseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MouseServiceServer will
// result in compilation errors.
type UnsafeMouseServiceServer interface {
	mustEmbedUnimplementedMouseServiceServer()
}

func RegisterMouseServiceServer(s grpc.ServiceRegistrar, srv MouseServiceServer) {
	s.RegisterService(&MouseService_ServiceDesc, srv)
}

func _MouseService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MouseServiceServer).Connect(&mouseServiceConnectServer{stream})
}

type MouseService_ConnectServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type mouseServiceConnectServer struct {
	grpc.ServerStream
}

func (x *mouseServiceConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mouseServiceConnectServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MouseService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MouseServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MouseService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MouseServiceServer).Disconnect(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MouseService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MouseServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MouseService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MouseServiceServer).Stat(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MouseService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MouseServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MouseService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MouseServiceServer).Start(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _MouseService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MouseServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MouseService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MouseServiceServer).Stop(ctx, req.(*StopTask))
	}
	return interceptor(ctx, in, info, handler)
}

// MouseService_ServiceDesc is the grpc.ServiceDesc for MouseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MouseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MouseService",
	HandlerType: (*MouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disconnect",
			Handler:    _MouseService_Disconnect_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _MouseService_Stat_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _MouseService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _MouseService_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _MouseService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mouse.proto",
}
