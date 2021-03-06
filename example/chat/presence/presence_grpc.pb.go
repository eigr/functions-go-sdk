// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package presence

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PresenceClient is the client API for Presence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresenceClient interface {
	// Connect the given user. They will stay connected as long as the stream stays open.
	Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (Presence_ConnectClient, error)
	// Monitor the online status of the given user.
	Monitor(ctx context.Context, in *User, opts ...grpc.CallOption) (Presence_MonitorClient, error)
}

type presenceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceClient(cc grpc.ClientConnInterface) PresenceClient {
	return &presenceClient{cc}
}

func (c *presenceClient) Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (Presence_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Presence_serviceDesc.Streams[0], "/functions.samples.chat.presence.Presence/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Presence_ConnectClient interface {
	Recv() (*Empty, error)
	grpc.ClientStream
}

type presenceConnectClient struct {
	grpc.ClientStream
}

func (x *presenceConnectClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *presenceClient) Monitor(ctx context.Context, in *User, opts ...grpc.CallOption) (Presence_MonitorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Presence_serviceDesc.Streams[1], "/functions.samples.chat.presence.Presence/Monitor", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenceMonitorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Presence_MonitorClient interface {
	Recv() (*OnlineStatus, error)
	grpc.ClientStream
}

type presenceMonitorClient struct {
	grpc.ClientStream
}

func (x *presenceMonitorClient) Recv() (*OnlineStatus, error) {
	m := new(OnlineStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PresenceServer is the server API for Presence service.
// All implementations must embed UnimplementedPresenceServer
// for forward compatibility
type PresenceServer interface {
	// Connect the given user. They will stay connected as long as the stream stays open.
	Connect(*User, Presence_ConnectServer) error
	// Monitor the online status of the given user.
	Monitor(*User, Presence_MonitorServer) error
	mustEmbedUnimplementedPresenceServer()
}

// UnimplementedPresenceServer must be embedded to have forward compatible implementations.
type UnimplementedPresenceServer struct {
}

func (UnimplementedPresenceServer) Connect(*User, Presence_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedPresenceServer) Monitor(*User, Presence_MonitorServer) error {
	return status.Errorf(codes.Unimplemented, "method Monitor not implemented")
}
func (UnimplementedPresenceServer) mustEmbedUnimplementedPresenceServer() {}

// UnsafePresenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresenceServer will
// result in compilation errors.
type UnsafePresenceServer interface {
	mustEmbedUnimplementedPresenceServer()
}

func RegisterPresenceServer(s grpc.ServiceRegistrar, srv PresenceServer) {
	s.RegisterService(&_Presence_serviceDesc, srv)
}

func _Presence_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenceServer).Connect(m, &presenceConnectServer{stream})
}

type Presence_ConnectServer interface {
	Send(*Empty) error
	grpc.ServerStream
}

type presenceConnectServer struct {
	grpc.ServerStream
}

func (x *presenceConnectServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _Presence_Monitor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenceServer).Monitor(m, &presenceMonitorServer{stream})
}

type Presence_MonitorServer interface {
	Send(*OnlineStatus) error
	grpc.ServerStream
}

type presenceMonitorServer struct {
	grpc.ServerStream
}

func (x *presenceMonitorServer) Send(m *OnlineStatus) error {
	return x.ServerStream.SendMsg(m)
}

var _Presence_serviceDesc = grpc.ServiceDesc{
	ServiceName: "functions.samples.chat.presence.Presence",
	HandlerType: (*PresenceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Presence_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Monitor",
			Handler:       _Presence_Monitor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "presence.proto",
}
