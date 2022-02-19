// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package eventlogeventing

import (
	context "context"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EventLogSubscriberModelClient is the client API for EventLogSubscriberModel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventLogSubscriberModelClient interface {
	ProcessEventOne(ctx context.Context, in *EventOne, opts ...grpc.CallOption) (*Response, error)
	ProcessEventTwo(ctx context.Context, in *EventTwo, opts ...grpc.CallOption) (EventLogSubscriberModel_ProcessEventTwoClient, error)
	Effect(ctx context.Context, in *EffectRequest, opts ...grpc.CallOption) (*Response, error)
	ProcessAnyEvent(ctx context.Context, in *any.Any, opts ...grpc.CallOption) (*Response, error)
}

type eventLogSubscriberModelClient struct {
	cc grpc.ClientConnInterface
}

func NewEventLogSubscriberModelClient(cc grpc.ClientConnInterface) EventLogSubscriberModelClient {
	return &eventLogSubscriberModelClient{cc}
}

func (c *eventLogSubscriberModelClient) ProcessEventOne(ctx context.Context, in *EventOne, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/ProcessEventOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventLogSubscriberModelClient) ProcessEventTwo(ctx context.Context, in *EventTwo, opts ...grpc.CallOption) (EventLogSubscriberModel_ProcessEventTwoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventLogSubscriberModel_serviceDesc.Streams[0], "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/ProcessEventTwo", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventLogSubscriberModelProcessEventTwoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventLogSubscriberModel_ProcessEventTwoClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type eventLogSubscriberModelProcessEventTwoClient struct {
	grpc.ClientStream
}

func (x *eventLogSubscriberModelProcessEventTwoClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventLogSubscriberModelClient) Effect(ctx context.Context, in *EffectRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/Effect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventLogSubscriberModelClient) ProcessAnyEvent(ctx context.Context, in *any.Any, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/ProcessAnyEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventLogSubscriberModelServer is the server API for EventLogSubscriberModel service.
// All implementations must embed UnimplementedEventLogSubscriberModelServer
// for forward compatibility
type EventLogSubscriberModelServer interface {
	ProcessEventOne(context.Context, *EventOne) (*Response, error)
	ProcessEventTwo(*EventTwo, EventLogSubscriberModel_ProcessEventTwoServer) error
	Effect(context.Context, *EffectRequest) (*Response, error)
	ProcessAnyEvent(context.Context, *any.Any) (*Response, error)
	mustEmbedUnimplementedEventLogSubscriberModelServer()
}

// UnimplementedEventLogSubscriberModelServer must be embedded to have forward compatible implementations.
type UnimplementedEventLogSubscriberModelServer struct {
}

func (UnimplementedEventLogSubscriberModelServer) ProcessEventOne(context.Context, *EventOne) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessEventOne not implemented")
}
func (UnimplementedEventLogSubscriberModelServer) ProcessEventTwo(*EventTwo, EventLogSubscriberModel_ProcessEventTwoServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessEventTwo not implemented")
}
func (UnimplementedEventLogSubscriberModelServer) Effect(context.Context, *EffectRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Effect not implemented")
}
func (UnimplementedEventLogSubscriberModelServer) ProcessAnyEvent(context.Context, *any.Any) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessAnyEvent not implemented")
}
func (UnimplementedEventLogSubscriberModelServer) mustEmbedUnimplementedEventLogSubscriberModelServer() {
}

// UnsafeEventLogSubscriberModelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventLogSubscriberModelServer will
// result in compilation errors.
type UnsafeEventLogSubscriberModelServer interface {
	mustEmbedUnimplementedEventLogSubscriberModelServer()
}

func RegisterEventLogSubscriberModelServer(s grpc.ServiceRegistrar, srv EventLogSubscriberModelServer) {
	s.RegisterService(&_EventLogSubscriberModel_serviceDesc, srv)
}

func _EventLogSubscriberModel_ProcessEventOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventOne)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventLogSubscriberModelServer).ProcessEventOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/ProcessEventOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventLogSubscriberModelServer).ProcessEventOne(ctx, req.(*EventOne))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventLogSubscriberModel_ProcessEventTwo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventTwo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventLogSubscriberModelServer).ProcessEventTwo(m, &eventLogSubscriberModelProcessEventTwoServer{stream})
}

type EventLogSubscriberModel_ProcessEventTwoServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type eventLogSubscriberModelProcessEventTwoServer struct {
	grpc.ServerStream
}

func (x *eventLogSubscriberModelProcessEventTwoServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _EventLogSubscriberModel_Effect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EffectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventLogSubscriberModelServer).Effect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/Effect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventLogSubscriberModelServer).Effect(ctx, req.(*EffectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventLogSubscriberModel_ProcessAnyEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(any.Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventLogSubscriberModelServer).ProcessAnyEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.tck.model.eventlogeventing.EventLogSubscriberModel/ProcessAnyEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventLogSubscriberModelServer).ProcessAnyEvent(ctx, req.(*any.Any))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventLogSubscriberModel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "functions.tck.model.eventlogeventing.EventLogSubscriberModel",
	HandlerType: (*EventLogSubscriberModelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessEventOne",
			Handler:    _EventLogSubscriberModel_ProcessEventOne_Handler,
		},
		{
			MethodName: "Effect",
			Handler:    _EventLogSubscriberModel_Effect_Handler,
		},
		{
			MethodName: "ProcessAnyEvent",
			Handler:    _EventLogSubscriberModel_ProcessAnyEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessEventTwo",
			Handler:       _EventLogSubscriberModel_ProcessEventTwo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eventlogeventing.proto",
}

// EventSourcedEntityOneClient is the client API for EventSourcedEntityOne service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedEntityOneClient interface {
	EmitEvent(ctx context.Context, in *EmitEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type eventSourcedEntityOneClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedEntityOneClient(cc grpc.ClientConnInterface) EventSourcedEntityOneClient {
	return &eventSourcedEntityOneClient{cc}
}

func (c *eventSourcedEntityOneClient) EmitEvent(ctx context.Context, in *EmitEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/functions.tck.model.eventlogeventing.EventSourcedEntityOne/EmitEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSourcedEntityOneServer is the server API for EventSourcedEntityOne service.
// All implementations must embed UnimplementedEventSourcedEntityOneServer
// for forward compatibility
type EventSourcedEntityOneServer interface {
	EmitEvent(context.Context, *EmitEventRequest) (*empty.Empty, error)
	mustEmbedUnimplementedEventSourcedEntityOneServer()
}

// UnimplementedEventSourcedEntityOneServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedEntityOneServer struct {
}

func (UnimplementedEventSourcedEntityOneServer) EmitEvent(context.Context, *EmitEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitEvent not implemented")
}
func (UnimplementedEventSourcedEntityOneServer) mustEmbedUnimplementedEventSourcedEntityOneServer() {}

// UnsafeEventSourcedEntityOneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedEntityOneServer will
// result in compilation errors.
type UnsafeEventSourcedEntityOneServer interface {
	mustEmbedUnimplementedEventSourcedEntityOneServer()
}

func RegisterEventSourcedEntityOneServer(s grpc.ServiceRegistrar, srv EventSourcedEntityOneServer) {
	s.RegisterService(&_EventSourcedEntityOne_serviceDesc, srv)
}

func _EventSourcedEntityOne_EmitEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSourcedEntityOneServer).EmitEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.tck.model.eventlogeventing.EventSourcedEntityOne/EmitEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSourcedEntityOneServer).EmitEvent(ctx, req.(*EmitEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventSourcedEntityOne_serviceDesc = grpc.ServiceDesc{
	ServiceName: "functions.tck.model.eventlogeventing.EventSourcedEntityOne",
	HandlerType: (*EventSourcedEntityOneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmitEvent",
			Handler:    _EventSourcedEntityOne_EmitEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventlogeventing.proto",
}

// EventSourcedEntityTwoClient is the client API for EventSourcedEntityTwo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedEntityTwoClient interface {
	EmitJsonEvent(ctx context.Context, in *JsonEvent, opts ...grpc.CallOption) (*empty.Empty, error)
}

type eventSourcedEntityTwoClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedEntityTwoClient(cc grpc.ClientConnInterface) EventSourcedEntityTwoClient {
	return &eventSourcedEntityTwoClient{cc}
}

func (c *eventSourcedEntityTwoClient) EmitJsonEvent(ctx context.Context, in *JsonEvent, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/functions.tck.model.eventlogeventing.EventSourcedEntityTwo/EmitJsonEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSourcedEntityTwoServer is the server API for EventSourcedEntityTwo service.
// All implementations must embed UnimplementedEventSourcedEntityTwoServer
// for forward compatibility
type EventSourcedEntityTwoServer interface {
	EmitJsonEvent(context.Context, *JsonEvent) (*empty.Empty, error)
	mustEmbedUnimplementedEventSourcedEntityTwoServer()
}

// UnimplementedEventSourcedEntityTwoServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedEntityTwoServer struct {
}

func (UnimplementedEventSourcedEntityTwoServer) EmitJsonEvent(context.Context, *JsonEvent) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitJsonEvent not implemented")
}
func (UnimplementedEventSourcedEntityTwoServer) mustEmbedUnimplementedEventSourcedEntityTwoServer() {}

// UnsafeEventSourcedEntityTwoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedEntityTwoServer will
// result in compilation errors.
type UnsafeEventSourcedEntityTwoServer interface {
	mustEmbedUnimplementedEventSourcedEntityTwoServer()
}

func RegisterEventSourcedEntityTwoServer(s grpc.ServiceRegistrar, srv EventSourcedEntityTwoServer) {
	s.RegisterService(&_EventSourcedEntityTwo_serviceDesc, srv)
}

func _EventSourcedEntityTwo_EmitJsonEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSourcedEntityTwoServer).EmitJsonEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/functions.tck.model.eventlogeventing.EventSourcedEntityTwo/EmitJsonEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSourcedEntityTwoServer).EmitJsonEvent(ctx, req.(*JsonEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventSourcedEntityTwo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "functions.tck.model.eventlogeventing.EventSourcedEntityTwo",
	HandlerType: (*EventSourcedEntityTwoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmitJsonEvent",
			Handler:    _EventSourcedEntityTwo_EmitJsonEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventlogeventing.proto",
}
