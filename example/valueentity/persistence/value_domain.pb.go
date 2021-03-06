// Copyright 2019 Lightbend Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// These are the messages that get persisted - the current state (Cart).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: value_domain.proto

package persistence

import (
	proto "github.com/golang/protobuf/proto"
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

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_value_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_value_domain_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *LineItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LineItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// The shopping cart state.
type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*LineItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_value_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_value_domain_proto_rawDescGZIP(), []int{1}
}

func (x *Cart) GetItems() []*LineItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_value_domain_proto protoreflect.FileDescriptor

var file_value_domain_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x58, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x50, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x67, 0x72, 0x2f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x3b, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_value_domain_proto_rawDescOnce sync.Once
	file_value_domain_proto_rawDescData = file_value_domain_proto_rawDesc
)

func file_value_domain_proto_rawDescGZIP() []byte {
	file_value_domain_proto_rawDescOnce.Do(func() {
		file_value_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_value_domain_proto_rawDescData)
	})
	return file_value_domain_proto_rawDescData
}

var file_value_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_value_domain_proto_goTypes = []interface{}{
	(*LineItem)(nil), // 0: com.example.valueentity.shoppingcart.persistence.LineItem
	(*Cart)(nil),     // 1: com.example.valueentity.shoppingcart.persistence.Cart
}
var file_value_domain_proto_depIdxs = []int32{
	0, // 0: com.example.valueentity.shoppingcart.persistence.Cart.items:type_name -> com.example.valueentity.shoppingcart.persistence.LineItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_value_domain_proto_init() }
func file_value_domain_proto_init() {
	if File_value_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_value_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
		file_value_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_value_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_value_domain_proto_goTypes,
		DependencyIndexes: file_value_domain_proto_depIdxs,
		MessageInfos:      file_value_domain_proto_msgTypes,
	}.Build()
	File_value_domain_proto = out.File
	file_value_domain_proto_rawDesc = nil
	file_value_domain_proto_goTypes = nil
	file_value_domain_proto_depIdxs = nil
}
