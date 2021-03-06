// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ActionProtocolClient is the client API for ActionProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionProtocolClient interface {
	// Handle a unary command.
	//
	// The input command will contain the service name, command name, request metadata and the command
	// payload. The reply may contain a direct reply, a forward or a failure, and it may contain many
	// side effects.
	HandleUnary(ctx context.Context, in *ActionCommand, opts ...grpc.CallOption) (*ActionResponse, error)
	// Handle a streamed in command.
	//
	// The first message in will contain the request metadata, including the service name and command
	// name. It will not have an associated payload set. This will be followed by zero to many messages
	// in with a payload, but no service name or command name set.
	//
	// If the underlying transport supports per stream metadata, rather than per message metadata, then
	// that metadata will only be included in the metadata of the first message. In contrast, if the
	// underlying transport supports per message metadata, there will be no metadata on the first message,
	// the metadata will instead be found on each subsequent message.
	//
	// The semantics of stream closure in this protocol map 1:1 with the semantics of gRPC stream closure,
	// that is, when the client closes the stream, the stream is considered half closed, and the server
	// should eventually, but not necessarily immediately, send a response message with a status code and
	// trailers.
	//
	// If however the server sends a response message before the client closes the stream, the stream is
	// completely closed, and the client should handle this and stop sending more messages.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamedIn(ctx context.Context, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedInClient, error)
	// Handle a streamed out command.
	//
	// The input command will contain the service name, command name, request metadata and the command
	// payload. Zero or more replies may be sent, each containing either a direct reply, a forward or a
	// failure, and each may contain many side effects. The stream to the client will be closed when the
	// this stream is closed, with the same status as this stream is closed with.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamedOut(ctx context.Context, in *ActionCommand, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedOutClient, error)
	// Handle a full duplex streamed command.
	//
	// The first message in will contain the request metadata, including the service name and command
	// name. It will not have an associated payload set. This will be followed by zero to many messages
	// in with a payload, but no service name or command name set.
	//
	// Zero or more replies may be sent, each containing either a direct reply, a forward or a failure,
	// and each may contain many side effects.
	//
	// If the underlying transport supports per stream metadata, rather than per message metadata, then
	// that metadata will only be included in the metadata of the first message. In contrast, if the
	// underlying transport supports per message metadata, there will be no metadata on the first message,
	// the metadata will instead be found on each subsequent message.
	//
	// The semantics of stream closure in this protocol map 1:1 with the semantics of gRPC stream closure,
	// that is, when the client closes the stream, the stream is considered half closed, and the server
	// should eventually, but not necessarily immediately, close the stream with a status code and
	// trailers.
	//
	// If however the server closes the stream with a status code and trailers, the stream is immediately
	// considered completely closed, and no further messages sent by the client will be handled by the
	// server.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamed(ctx context.Context, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedClient, error)
}

type actionProtocolClient struct {
	cc grpc.ClientConnInterface
}

func NewActionProtocolClient(cc grpc.ClientConnInterface) ActionProtocolClient {
	return &actionProtocolClient{cc}
}

func (c *actionProtocolClient) HandleUnary(ctx context.Context, in *ActionCommand, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/functions.action.ActionProtocol/handleUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionProtocolClient) HandleStreamedIn(ctx context.Context, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedInClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActionProtocol_serviceDesc.Streams[0], "/functions.action.ActionProtocol/handleStreamedIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &actionProtocolHandleStreamedInClient{stream}
	return x, nil
}

type ActionProtocol_HandleStreamedInClient interface {
	Send(*ActionCommand) error
	CloseAndRecv() (*ActionResponse, error)
	grpc.ClientStream
}

type actionProtocolHandleStreamedInClient struct {
	grpc.ClientStream
}

func (x *actionProtocolHandleStreamedInClient) Send(m *ActionCommand) error {
	return x.ClientStream.SendMsg(m)
}

func (x *actionProtocolHandleStreamedInClient) CloseAndRecv() (*ActionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *actionProtocolClient) HandleStreamedOut(ctx context.Context, in *ActionCommand, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedOutClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActionProtocol_serviceDesc.Streams[1], "/functions.action.ActionProtocol/handleStreamedOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &actionProtocolHandleStreamedOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActionProtocol_HandleStreamedOutClient interface {
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type actionProtocolHandleStreamedOutClient struct {
	grpc.ClientStream
}

func (x *actionProtocolHandleStreamedOutClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *actionProtocolClient) HandleStreamed(ctx context.Context, opts ...grpc.CallOption) (ActionProtocol_HandleStreamedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ActionProtocol_serviceDesc.Streams[2], "/functions.action.ActionProtocol/handleStreamed", opts...)
	if err != nil {
		return nil, err
	}
	x := &actionProtocolHandleStreamedClient{stream}
	return x, nil
}

type ActionProtocol_HandleStreamedClient interface {
	Send(*ActionCommand) error
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type actionProtocolHandleStreamedClient struct {
	grpc.ClientStream
}

func (x *actionProtocolHandleStreamedClient) Send(m *ActionCommand) error {
	return x.ClientStream.SendMsg(m)
}

func (x *actionProtocolHandleStreamedClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActionProtocolServer is the server API for ActionProtocol service.
// All implementations must embed UnimplementedActionProtocolServer
// for forward compatibility
type ActionProtocolServer interface {
	// Handle a unary command.
	//
	// The input command will contain the service name, command name, request metadata and the command
	// payload. The reply may contain a direct reply, a forward or a failure, and it may contain many
	// side effects.
	HandleUnary(context.Context, *ActionCommand) (*ActionResponse, error)
	// Handle a streamed in command.
	//
	// The first message in will contain the request metadata, including the service name and command
	// name. It will not have an associated payload set. This will be followed by zero to many messages
	// in with a payload, but no service name or command name set.
	//
	// If the underlying transport supports per stream metadata, rather than per message metadata, then
	// that metadata will only be included in the metadata of the first message. In contrast, if the
	// underlying transport supports per message metadata, there will be no metadata on the first message,
	// the metadata will instead be found on each subsequent message.
	//
	// The semantics of stream closure in this protocol map 1:1 with the semantics of gRPC stream closure,
	// that is, when the client closes the stream, the stream is considered half closed, and the server
	// should eventually, but not necessarily immediately, send a response message with a status code and
	// trailers.
	//
	// If however the server sends a response message before the client closes the stream, the stream is
	// completely closed, and the client should handle this and stop sending more messages.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamedIn(ActionProtocol_HandleStreamedInServer) error
	// Handle a streamed out command.
	//
	// The input command will contain the service name, command name, request metadata and the command
	// payload. Zero or more replies may be sent, each containing either a direct reply, a forward or a
	// failure, and each may contain many side effects. The stream to the client will be closed when the
	// this stream is closed, with the same status as this stream is closed with.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamedOut(*ActionCommand, ActionProtocol_HandleStreamedOutServer) error
	// Handle a full duplex streamed command.
	//
	// The first message in will contain the request metadata, including the service name and command
	// name. It will not have an associated payload set. This will be followed by zero to many messages
	// in with a payload, but no service name or command name set.
	//
	// Zero or more replies may be sent, each containing either a direct reply, a forward or a failure,
	// and each may contain many side effects.
	//
	// If the underlying transport supports per stream metadata, rather than per message metadata, then
	// that metadata will only be included in the metadata of the first message. In contrast, if the
	// underlying transport supports per message metadata, there will be no metadata on the first message,
	// the metadata will instead be found on each subsequent message.
	//
	// The semantics of stream closure in this protocol map 1:1 with the semantics of gRPC stream closure,
	// that is, when the client closes the stream, the stream is considered half closed, and the server
	// should eventually, but not necessarily immediately, close the stream with a status code and
	// trailers.
	//
	// If however the server closes the stream with a status code and trailers, the stream is immediately
	// considered completely closed, and no further messages sent by the client will be handled by the
	// server.
	//
	// Either the client or the server may cancel the stream at any time, cancellation is indicated
	// through an HTTP2 stream RST message.
	HandleStreamed(ActionProtocol_HandleStreamedServer) error
	mustEmbedUnimplementedActionProtocolServer()
}

// UnimplementedActionProtocolServer must be embedded to have forward compatible implementations.
type UnimplementedActionProtocolServer struct {
}

func (UnimplementedActionProtocolServer) HandleUnary(context.Context, *ActionCommand) (*ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUnary not implemented")
}
func (UnimplementedActionProtocolServer) HandleStreamedIn(ActionProtocol_HandleStreamedInServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleStreamedIn not implemented")
}
func (UnimplementedActionProtocolServer) HandleStreamedOut(*ActionCommand, ActionProtocol_HandleStreamedOutServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleStreamedOut not implemented")
}
func (UnimplementedActionProtocolServer) HandleStreamed(ActionProtocol_HandleStreamedServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleStreamed not implemented")
}
func (UnimplementedActionProtocolServer) mustEmbedUnimplementedActionProtocolServer() {}

// UnsafeActionProtocolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionProtocolServer will
// result in compilation errors.
type UnsafeActionProtocolServer interface {
	mustEmbedUnimplementedActionProtocolServer()
}

func RegisterActionProtocolServer(s grpc.ServiceRegistrar, srv ActionProtocolServer) {
	s.RegisterService(&_ActionProtocol_serviceDesc, srv)
}

func _ActionProtocol_HandleUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionProtocolServer).HandleUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.action.ActionProtocol/handleUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionProtocolServer).HandleUnary(ctx, req.(*ActionCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionProtocol_HandleStreamedIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ActionProtocolServer).HandleStreamedIn(&actionProtocolHandleStreamedInServer{stream})
}

type ActionProtocol_HandleStreamedInServer interface {
	SendAndClose(*ActionResponse) error
	Recv() (*ActionCommand, error)
	grpc.ServerStream
}

type actionProtocolHandleStreamedInServer struct {
	grpc.ServerStream
}

func (x *actionProtocolHandleStreamedInServer) SendAndClose(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *actionProtocolHandleStreamedInServer) Recv() (*ActionCommand, error) {
	m := new(ActionCommand)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ActionProtocol_HandleStreamedOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActionCommand)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionProtocolServer).HandleStreamedOut(m, &actionProtocolHandleStreamedOutServer{stream})
}

type ActionProtocol_HandleStreamedOutServer interface {
	Send(*ActionResponse) error
	grpc.ServerStream
}

type actionProtocolHandleStreamedOutServer struct {
	grpc.ServerStream
}

func (x *actionProtocolHandleStreamedOutServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ActionProtocol_HandleStreamed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ActionProtocolServer).HandleStreamed(&actionProtocolHandleStreamedServer{stream})
}

type ActionProtocol_HandleStreamedServer interface {
	Send(*ActionResponse) error
	Recv() (*ActionCommand, error)
	grpc.ServerStream
}

type actionProtocolHandleStreamedServer struct {
	grpc.ServerStream
}

func (x *actionProtocolHandleStreamedServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *actionProtocolHandleStreamedServer) Recv() (*ActionCommand, error) {
	m := new(ActionCommand)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ActionProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "functions.action.ActionProtocol",
	HandlerType: (*ActionProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handleUnary",
			Handler:    _ActionProtocol_HandleUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "handleStreamedIn",
			Handler:       _ActionProtocol_HandleStreamedIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "handleStreamedOut",
			Handler:       _ActionProtocol_HandleStreamedOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "handleStreamed",
			Handler:       _ActionProtocol_HandleStreamed_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "action.proto",
}
