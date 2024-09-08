// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package book

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

// BookGapiClient is the client API for BookGapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookGapiClient interface {
	InsertBook(ctx context.Context, in *InsertBookRequest, opts ...grpc.CallOption) (*InsertBookResponse, error)
	GetListBooks(ctx context.Context, in *GetListBooksRequest, opts ...grpc.CallOption) (*GetListBooksResponse, error)
	GetBookById(ctx context.Context, in *GetBookByIdRequest, opts ...grpc.CallOption) (*GetBookByIdResponse, error)
	UpdateBookById(ctx context.Context, in *UpdateBookByIdRequest, opts ...grpc.CallOption) (*UpdateBookByIdResponse, error)
	DeleteBookById(ctx context.Context, in *DeleteBookByIdRequest, opts ...grpc.CallOption) (*DeleteBookByIdResponse, error)
}

type bookGapiClient struct {
	cc grpc.ClientConnInterface
}

func NewBookGapiClient(cc grpc.ClientConnInterface) BookGapiClient {
	return &bookGapiClient{cc}
}

func (c *bookGapiClient) InsertBook(ctx context.Context, in *InsertBookRequest, opts ...grpc.CallOption) (*InsertBookResponse, error) {
	out := new(InsertBookResponse)
	err := c.cc.Invoke(ctx, "/book.BookGapi/InsertBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookGapiClient) GetListBooks(ctx context.Context, in *GetListBooksRequest, opts ...grpc.CallOption) (*GetListBooksResponse, error) {
	out := new(GetListBooksResponse)
	err := c.cc.Invoke(ctx, "/book.BookGapi/GetListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookGapiClient) GetBookById(ctx context.Context, in *GetBookByIdRequest, opts ...grpc.CallOption) (*GetBookByIdResponse, error) {
	out := new(GetBookByIdResponse)
	err := c.cc.Invoke(ctx, "/book.BookGapi/GetBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookGapiClient) UpdateBookById(ctx context.Context, in *UpdateBookByIdRequest, opts ...grpc.CallOption) (*UpdateBookByIdResponse, error) {
	out := new(UpdateBookByIdResponse)
	err := c.cc.Invoke(ctx, "/book.BookGapi/UpdateBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookGapiClient) DeleteBookById(ctx context.Context, in *DeleteBookByIdRequest, opts ...grpc.CallOption) (*DeleteBookByIdResponse, error) {
	out := new(DeleteBookByIdResponse)
	err := c.cc.Invoke(ctx, "/book.BookGapi/DeleteBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookGapiServer is the server API for BookGapi service.
// All implementations must embed UnimplementedBookGapiServer
// for forward compatibility
type BookGapiServer interface {
	InsertBook(context.Context, *InsertBookRequest) (*InsertBookResponse, error)
	GetListBooks(context.Context, *GetListBooksRequest) (*GetListBooksResponse, error)
	GetBookById(context.Context, *GetBookByIdRequest) (*GetBookByIdResponse, error)
	UpdateBookById(context.Context, *UpdateBookByIdRequest) (*UpdateBookByIdResponse, error)
	DeleteBookById(context.Context, *DeleteBookByIdRequest) (*DeleteBookByIdResponse, error)
	mustEmbedUnimplementedBookGapiServer()
}

// UnimplementedBookGapiServer must be embedded to have forward compatible implementations.
type UnimplementedBookGapiServer struct {
}

func (UnimplementedBookGapiServer) InsertBook(context.Context, *InsertBookRequest) (*InsertBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
}
func (UnimplementedBookGapiServer) GetListBooks(context.Context, *GetListBooksRequest) (*GetListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListBooks not implemented")
}
func (UnimplementedBookGapiServer) GetBookById(context.Context, *GetBookByIdRequest) (*GetBookByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookById not implemented")
}
func (UnimplementedBookGapiServer) UpdateBookById(context.Context, *UpdateBookByIdRequest) (*UpdateBookByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookById not implemented")
}
func (UnimplementedBookGapiServer) DeleteBookById(context.Context, *DeleteBookByIdRequest) (*DeleteBookByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookById not implemented")
}
func (UnimplementedBookGapiServer) mustEmbedUnimplementedBookGapiServer() {}

// UnsafeBookGapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookGapiServer will
// result in compilation errors.
type UnsafeBookGapiServer interface {
	mustEmbedUnimplementedBookGapiServer()
}

func RegisterBookGapiServer(s grpc.ServiceRegistrar, srv BookGapiServer) {
	s.RegisterService(&BookGapi_ServiceDesc, srv)
}

func _BookGapi_InsertBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGapiServer).InsertBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookGapi/InsertBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGapiServer).InsertBook(ctx, req.(*InsertBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookGapi_GetListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGapiServer).GetListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookGapi/GetListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGapiServer).GetListBooks(ctx, req.(*GetListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookGapi_GetBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGapiServer).GetBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookGapi/GetBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGapiServer).GetBookById(ctx, req.(*GetBookByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookGapi_UpdateBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGapiServer).UpdateBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookGapi/UpdateBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGapiServer).UpdateBookById(ctx, req.(*UpdateBookByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookGapi_DeleteBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGapiServer).DeleteBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookGapi/DeleteBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGapiServer).DeleteBookById(ctx, req.(*DeleteBookByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookGapi_ServiceDesc is the grpc.ServiceDesc for BookGapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookGapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookGapi",
	HandlerType: (*BookGapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBook",
			Handler:    _BookGapi_InsertBook_Handler,
		},
		{
			MethodName: "GetListBooks",
			Handler:    _BookGapi_GetListBooks_Handler,
		},
		{
			MethodName: "GetBookById",
			Handler:    _BookGapi_GetBookById_Handler,
		},
		{
			MethodName: "UpdateBookById",
			Handler:    _BookGapi_UpdateBookById_Handler,
		},
		{
			MethodName: "DeleteBookById",
			Handler:    _BookGapi_DeleteBookById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book/book.proto",
}
