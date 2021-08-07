// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package valueentity

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ShoppingCartClient is the client API for ShoppingCart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShoppingCartClient interface {
	AddItem(ctx context.Context, in *AddLineItem, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveItem(ctx context.Context, in *RemoveLineItem, opts ...grpc.CallOption) (*empty.Empty, error)
	GetCart(ctx context.Context, in *GetShoppingCart, opts ...grpc.CallOption) (*Cart, error)
	RemoveCart(ctx context.Context, in *RemoveShoppingCart, opts ...grpc.CallOption) (*empty.Empty, error)
}

type shoppingCartClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingCartClient(cc grpc.ClientConnInterface) ShoppingCartClient {
	return &shoppingCartClient{cc}
}

func (c *shoppingCartClient) AddItem(ctx context.Context, in *AddLineItem, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.example.valueentity.shoppingcart.ShoppingCart/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartClient) RemoveItem(ctx context.Context, in *RemoveLineItem, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.example.valueentity.shoppingcart.ShoppingCart/RemoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartClient) GetCart(ctx context.Context, in *GetShoppingCart, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/com.example.valueentity.shoppingcart.ShoppingCart/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartClient) RemoveCart(ctx context.Context, in *RemoveShoppingCart, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.example.valueentity.shoppingcart.ShoppingCart/RemoveCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingCartServer is the server API for ShoppingCart service.
// All implementations must embed UnimplementedShoppingCartServer
// for forward compatibility
type ShoppingCartServer interface {
	AddItem(context.Context, *AddLineItem) (*empty.Empty, error)
	RemoveItem(context.Context, *RemoveLineItem) (*empty.Empty, error)
	GetCart(context.Context, *GetShoppingCart) (*Cart, error)
	RemoveCart(context.Context, *RemoveShoppingCart) (*empty.Empty, error)
	mustEmbedUnimplementedShoppingCartServer()
}

// UnimplementedShoppingCartServer must be embedded to have forward compatible implementations.
type UnimplementedShoppingCartServer struct {
}

func (UnimplementedShoppingCartServer) AddItem(context.Context, *AddLineItem) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedShoppingCartServer) RemoveItem(context.Context, *RemoveLineItem) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedShoppingCartServer) GetCart(context.Context, *GetShoppingCart) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedShoppingCartServer) RemoveCart(context.Context, *RemoveShoppingCart) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCart not implemented")
}
func (UnimplementedShoppingCartServer) mustEmbedUnimplementedShoppingCartServer() {}

// UnsafeShoppingCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShoppingCartServer will
// result in compilation errors.
type UnsafeShoppingCartServer interface {
	mustEmbedUnimplementedShoppingCartServer()
}

func RegisterShoppingCartServer(s grpc.ServiceRegistrar, srv ShoppingCartServer) {
	s.RegisterService(&_ShoppingCart_serviceDesc, srv)
}

func _ShoppingCart_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLineItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.valueentity.shoppingcart.ShoppingCart/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServer).AddItem(ctx, req.(*AddLineItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCart_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLineItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.valueentity.shoppingcart.ShoppingCart/RemoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServer).RemoveItem(ctx, req.(*RemoveLineItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCart_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingCart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.valueentity.shoppingcart.ShoppingCart/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServer).GetCart(ctx, req.(*GetShoppingCart))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCart_RemoveCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShoppingCart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServer).RemoveCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.valueentity.shoppingcart.ShoppingCart/RemoveCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServer).RemoveCart(ctx, req.(*RemoveShoppingCart))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingCart_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.example.valueentity.shoppingcart.ShoppingCart",
	HandlerType: (*ShoppingCartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ShoppingCart_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _ShoppingCart_RemoveItem_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _ShoppingCart_GetCart_Handler,
		},
		{
			MethodName: "RemoveCart",
			Handler:    _ShoppingCart_RemoveCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "value_shoppingcart.proto",
}