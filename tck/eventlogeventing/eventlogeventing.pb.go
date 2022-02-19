// Copyright 2019 Lightbend Inc.
// Copyright 2021 The eigr.io Authors.
// Use of this source code is governed by the Apache License
// that can be found in the LICENSE file.

//
// == Cloudstate TCK model test for event log eventing ==
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: eventlogeventing.proto

package eventlogeventing

import (
	_ "github.com/eigr/functions-go-sdk/functions"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
// An `EmitEventRequest` is received by the `EventSourcedEntityOne` entity to instruct it to emit either an `EventOne`
// or an `EventTwo`.
//
type EmitEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Event:
	//	*EmitEventRequest_EventOne
	//	*EmitEventRequest_EventTwo
	Event isEmitEventRequest_Event `protobuf_oneof:"event"`
}

func (x *EmitEventRequest) Reset() {
	*x = EmitEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitEventRequest) ProtoMessage() {}

func (x *EmitEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitEventRequest.ProtoReflect.Descriptor instead.
func (*EmitEventRequest) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{0}
}

func (x *EmitEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *EmitEventRequest) GetEvent() isEmitEventRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *EmitEventRequest) GetEventOne() *EventOne {
	if x, ok := x.GetEvent().(*EmitEventRequest_EventOne); ok {
		return x.EventOne
	}
	return nil
}

func (x *EmitEventRequest) GetEventTwo() *EventTwo {
	if x, ok := x.GetEvent().(*EmitEventRequest_EventTwo); ok {
		return x.EventTwo
	}
	return nil
}

type isEmitEventRequest_Event interface {
	isEmitEventRequest_Event()
}

type EmitEventRequest_EventOne struct {
	EventOne *EventOne `protobuf:"bytes,2,opt,name=event_one,json=eventOne,proto3,oneof"`
}

type EmitEventRequest_EventTwo struct {
	EventTwo *EventTwo `protobuf:"bytes,3,opt,name=event_two,json=eventTwo,proto3,oneof"`
}

func (*EmitEventRequest_EventOne) isEmitEventRequest_Event() {}

func (*EmitEventRequest_EventTwo) isEmitEventRequest_Event() {}

//
// An `EventOne` is an event emitted by the `EventSourcedEntityOne` entity and subscribed to by
// `EventLogSubscriberModel`.
//
type EventOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *ProcessStep `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *EventOne) Reset() {
	*x = EventOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventOne) ProtoMessage() {}

func (x *EventOne) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventOne.ProtoReflect.Descriptor instead.
func (*EventOne) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{1}
}

func (x *EventOne) GetStep() *ProcessStep {
	if x != nil {
		return x.Step
	}
	return nil
}

//
// An `EventTwo` is an event emitted by the `EventSourcedEntityOne` entity and subscribed to by
// `EventLogSubscriberModel`.
//
type EventTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step []*ProcessStep `protobuf:"bytes,2,rep,name=step,proto3" json:"step,omitempty"`
}

func (x *EventTwo) Reset() {
	*x = EventTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTwo) ProtoMessage() {}

func (x *EventTwo) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTwo.ProtoReflect.Descriptor instead.
func (*EventTwo) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{2}
}

func (x *EventTwo) GetStep() []*ProcessStep {
	if x != nil {
		return x.Step
	}
	return nil
}

//
// A `JsonEvent` is an event emitted by the `EventSourcedEntityTwo` entity and subscribed to by
// `EventLogSubscriberModel`.
//
type JsonEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JsonEvent) Reset() {
	*x = JsonEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonEvent) ProtoMessage() {}

func (x *JsonEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonEvent.ProtoReflect.Descriptor instead.
func (*JsonEvent) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{3}
}

func (x *JsonEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JsonEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// Each `ProcessStep` is one of:
//
// - Reply: reply with the given message in a `Response`.
// - Forward: forward to another service, in place of replying with a `Response`.
//
type ProcessStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Step:
	//	*ProcessStep_Reply
	//	*ProcessStep_Forward
	Step isProcessStep_Step `protobuf_oneof:"step"`
}

func (x *ProcessStep) Reset() {
	*x = ProcessStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStep) ProtoMessage() {}

func (x *ProcessStep) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStep.ProtoReflect.Descriptor instead.
func (*ProcessStep) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{4}
}

func (m *ProcessStep) GetStep() isProcessStep_Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (x *ProcessStep) GetReply() *Reply {
	if x, ok := x.GetStep().(*ProcessStep_Reply); ok {
		return x.Reply
	}
	return nil
}

func (x *ProcessStep) GetForward() *Forward {
	if x, ok := x.GetStep().(*ProcessStep_Forward); ok {
		return x.Forward
	}
	return nil
}

type isProcessStep_Step interface {
	isProcessStep_Step()
}

type ProcessStep_Reply struct {
	Reply *Reply `protobuf:"bytes,1,opt,name=reply,proto3,oneof"`
}

type ProcessStep_Forward struct {
	Forward *Forward `protobuf:"bytes,2,opt,name=forward,proto3,oneof"`
}

func (*ProcessStep_Reply) isProcessStep_Step() {}

func (*ProcessStep_Forward) isProcessStep_Step() {}

//
// Reply with a message in the response.
//
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{5}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// Replace the response with a forward to `functions.tck.model.eventlogeventing.EventLogSubscriberModel/Effect`.
// The payload must be an `EffectRequest` message with the given `message`.
//
type Forward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Forward) Reset() {
	*x = Forward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forward) ProtoMessage() {}

func (x *Forward) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forward.ProtoReflect.Descriptor instead.
func (*Forward) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{6}
}

func (x *Forward) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// The `Response` message must contain the message from the corresponding reply step.
//
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The `EffectRequest` message must contain the id from the SideEffect or Forward.
type EffectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EffectRequest) Reset() {
	*x = EffectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventlogeventing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectRequest) ProtoMessage() {}

func (x *EffectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventlogeventing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectRequest.ProtoReflect.Descriptor instead.
func (*EffectRequest) Descriptor() ([]byte, []int) {
	return file_eventlogeventing_proto_rawDescGZIP(), []int{8}
}

func (x *EffectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EffectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_eventlogeventing_proto protoreflect.FileDescriptor

var file_eventlogeventing_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd0, 0x01, 0x0a, 0x10, 0x45, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x48, 0x00, 0x52,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x52, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x46,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70,
	0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x52, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x77, 0x6f, 0x12, 0x46, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63,
	0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x3a, 0x0a, 0x09, 0x4a, 0x73,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x12, 0x44, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x07,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x22, 0x21, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39,
	0x0a, 0x0d, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa9, 0x04, 0x0a, 0x17, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x90, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xca, 0x43, 0x18,
	0x0a, 0x16, 0x1a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2d, 0x6f, 0x6e, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x2f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x1a, 0x2f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0xca, 0x43, 0x18, 0x0a, 0x16, 0x1a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x6f, 0x6e, 0x65, 0x30, 0x01, 0x12, 0x6f, 0x0a,
	0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6e, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xca, 0x43, 0x18, 0x0a, 0x16, 0x1a,
	0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2d, 0x74, 0x77, 0x6f, 0x32, 0x75, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x65, 0x12, 0x5c,
	0x0a, 0x09, 0x45, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x72, 0x0a, 0x15,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x77, 0x6f, 0x12, 0x59, 0x0a, 0x0d, 0x45, 0x6d, 0x69, 0x74, 0x4a, 0x73, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4a,
	0x73, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x61, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x67, 0x72, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x63,
	0x6b, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventlogeventing_proto_rawDescOnce sync.Once
	file_eventlogeventing_proto_rawDescData = file_eventlogeventing_proto_rawDesc
)

func file_eventlogeventing_proto_rawDescGZIP() []byte {
	file_eventlogeventing_proto_rawDescOnce.Do(func() {
		file_eventlogeventing_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventlogeventing_proto_rawDescData)
	})
	return file_eventlogeventing_proto_rawDescData
}

var file_eventlogeventing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_eventlogeventing_proto_goTypes = []interface{}{
	(*EmitEventRequest)(nil), // 0: functions.tck.model.eventlogeventing.EmitEventRequest
	(*EventOne)(nil),         // 1: functions.tck.model.eventlogeventing.EventOne
	(*EventTwo)(nil),         // 2: functions.tck.model.eventlogeventing.EventTwo
	(*JsonEvent)(nil),        // 3: functions.tck.model.eventlogeventing.JsonEvent
	(*ProcessStep)(nil),      // 4: functions.tck.model.eventlogeventing.ProcessStep
	(*Reply)(nil),            // 5: functions.tck.model.eventlogeventing.Reply
	(*Forward)(nil),          // 6: functions.tck.model.eventlogeventing.Forward
	(*Response)(nil),         // 7: functions.tck.model.eventlogeventing.Response
	(*EffectRequest)(nil),    // 8: functions.tck.model.eventlogeventing.EffectRequest
	(*any.Any)(nil),          // 9: google.protobuf.Any
	(*empty.Empty)(nil),      // 10: google.protobuf.Empty
}
var file_eventlogeventing_proto_depIdxs = []int32{
	1,  // 0: functions.tck.model.eventlogeventing.EmitEventRequest.event_one:type_name -> functions.tck.model.eventlogeventing.EventOne
	2,  // 1: functions.tck.model.eventlogeventing.EmitEventRequest.event_two:type_name -> functions.tck.model.eventlogeventing.EventTwo
	4,  // 2: functions.tck.model.eventlogeventing.EventOne.step:type_name -> functions.tck.model.eventlogeventing.ProcessStep
	4,  // 3: functions.tck.model.eventlogeventing.EventTwo.step:type_name -> functions.tck.model.eventlogeventing.ProcessStep
	5,  // 4: functions.tck.model.eventlogeventing.ProcessStep.reply:type_name -> functions.tck.model.eventlogeventing.Reply
	6,  // 5: functions.tck.model.eventlogeventing.ProcessStep.forward:type_name -> functions.tck.model.eventlogeventing.Forward
	1,  // 6: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessEventOne:input_type -> functions.tck.model.eventlogeventing.EventOne
	2,  // 7: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessEventTwo:input_type -> functions.tck.model.eventlogeventing.EventTwo
	8,  // 8: functions.tck.model.eventlogeventing.EventLogSubscriberModel.Effect:input_type -> functions.tck.model.eventlogeventing.EffectRequest
	9,  // 9: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessAnyEvent:input_type -> google.protobuf.Any
	0,  // 10: functions.tck.model.eventlogeventing.EventSourcedEntityOne.EmitEvent:input_type -> functions.tck.model.eventlogeventing.EmitEventRequest
	3,  // 11: functions.tck.model.eventlogeventing.EventSourcedEntityTwo.EmitJsonEvent:input_type -> functions.tck.model.eventlogeventing.JsonEvent
	7,  // 12: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessEventOne:output_type -> functions.tck.model.eventlogeventing.Response
	7,  // 13: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessEventTwo:output_type -> functions.tck.model.eventlogeventing.Response
	7,  // 14: functions.tck.model.eventlogeventing.EventLogSubscriberModel.Effect:output_type -> functions.tck.model.eventlogeventing.Response
	7,  // 15: functions.tck.model.eventlogeventing.EventLogSubscriberModel.ProcessAnyEvent:output_type -> functions.tck.model.eventlogeventing.Response
	10, // 16: functions.tck.model.eventlogeventing.EventSourcedEntityOne.EmitEvent:output_type -> google.protobuf.Empty
	10, // 17: functions.tck.model.eventlogeventing.EventSourcedEntityTwo.EmitJsonEvent:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_eventlogeventing_proto_init() }
func file_eventlogeventing_proto_init() {
	if File_eventlogeventing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventlogeventing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmitEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventOne); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTwo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventlogeventing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_eventlogeventing_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EmitEventRequest_EventOne)(nil),
		(*EmitEventRequest_EventTwo)(nil),
	}
	file_eventlogeventing_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ProcessStep_Reply)(nil),
		(*ProcessStep_Forward)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eventlogeventing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_eventlogeventing_proto_goTypes,
		DependencyIndexes: file_eventlogeventing_proto_depIdxs,
		MessageInfos:      file_eventlogeventing_proto_msgTypes,
	}.Build()
	File_eventlogeventing_proto = out.File
	file_eventlogeventing_proto_rawDesc = nil
	file_eventlogeventing_proto_goTypes = nil
	file_eventlogeventing_proto_depIdxs = nil
}
