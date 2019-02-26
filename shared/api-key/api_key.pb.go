// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api_key.proto

package api_key

import (
	context "context"
	fmt "fmt"
	v1 "github.com/confluentinc/ccloudapis/auth/v1"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("api_key.proto", fileDescriptor_3d0a7164f3256520) }

var fileDescriptor_3d0a7164f3256520 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0x8c,
	0xcf, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x4c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0xf3, 0xd2, 0x72, 0x4a,
	0x53, 0xf3, 0x4a, 0x32, 0xf3, 0x92, 0xf5, 0x93, 0x93, 0x73, 0xf2, 0x4b, 0x53, 0x12, 0x0b, 0x32,
	0x8b, 0xf5, 0x13, 0x4b, 0x4b, 0x32, 0xf4, 0xcb, 0x0c, 0xc1, 0x34, 0x44, 0xbf, 0xd1, 0x34, 0x26,
	0x2e, 0x36, 0xc7, 0x82, 0x4c, 0xef, 0xd4, 0x4a, 0xa1, 0x28, 0x2e, 0x36, 0xe7, 0xa2, 0xd4, 0xc4,
	0x92, 0x54, 0x21, 0x75, 0x3d, 0x84, 0x26, 0x3d, 0xb0, 0xe2, 0x32, 0x43, 0x3d, 0x88, 0x1c, 0x44,
	0x6d, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x2a, 0x61, 0x85, 0x05, 0x39, 0x95, 0x4a,
	0x0c, 0x20, 0xb3, 0x5d, 0x52, 0x73, 0x52, 0x71, 0x99, 0x0d, 0x91, 0x23, 0xc2, 0x6c, 0x54, 0x85,
	0x10, 0xb3, 0x43, 0xb8, 0x58, 0x7c, 0x32, 0x8b, 0x4b, 0x84, 0xb0, 0x6a, 0x70, 0x4f, 0x2d, 0x81,
	0xa8, 0x2e, 0x86, 0x99, 0xab, 0x4c, 0x48, 0x19, 0xd8, 0xd4, 0x24, 0x36, 0x70, 0xf8, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xac, 0x2d, 0xa9, 0xbe, 0x70, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiKeyClient is the client API for ApiKey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiKeyClient interface {
	Create(ctx context.Context, in *v1.CreateApiKeyRequest, opts ...grpc.CallOption) (*v1.CreateApiKeyReply, error)
	Delete(ctx context.Context, in *v1.DeleteApiKeyRequest, opts ...grpc.CallOption) (*v1.DeleteApiKeyReply, error)
	List(ctx context.Context, in *v1.GetApiKeysRequest, opts ...grpc.CallOption) (*v1.GetApiKeysReply, error)
}

type apiKeyClient struct {
	cc *grpc.ClientConn
}

func NewApiKeyClient(cc *grpc.ClientConn) ApiKeyClient {
	return &apiKeyClient{cc}
}

func (c *apiKeyClient) Create(ctx context.Context, in *v1.CreateApiKeyRequest, opts ...grpc.CallOption) (*v1.CreateApiKeyReply, error) {
	out := new(v1.CreateApiKeyReply)
	err := c.cc.Invoke(ctx, "/api_key.ApiKey/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyClient) Delete(ctx context.Context, in *v1.DeleteApiKeyRequest, opts ...grpc.CallOption) (*v1.DeleteApiKeyReply, error) {
	out := new(v1.DeleteApiKeyReply)
	err := c.cc.Invoke(ctx, "/api_key.ApiKey/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyClient) List(ctx context.Context, in *v1.GetApiKeysRequest, opts ...grpc.CallOption) (*v1.GetApiKeysReply, error) {
	out := new(v1.GetApiKeysReply)
	err := c.cc.Invoke(ctx, "/api_key.ApiKey/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiKeyServer is the server API for ApiKey service.
type ApiKeyServer interface {
	Create(context.Context, *v1.CreateApiKeyRequest) (*v1.CreateApiKeyReply, error)
	Delete(context.Context, *v1.DeleteApiKeyRequest) (*v1.DeleteApiKeyReply, error)
	List(context.Context, *v1.GetApiKeysRequest) (*v1.GetApiKeysReply, error)
}

func RegisterApiKeyServer(s *grpc.Server, srv ApiKeyServer) {
	s.RegisterService(&_ApiKey_serviceDesc, srv)
}

func _ApiKey_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_key.ApiKey/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServer).Create(ctx, req.(*v1.CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKey_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_key.ApiKey/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServer).Delete(ctx, req.(*v1.DeleteApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKey_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetApiKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_key.ApiKey/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServer).List(ctx, req.(*v1.GetApiKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiKey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api_key.ApiKey",
	HandlerType: (*ApiKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApiKey_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApiKey_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApiKey_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_key.proto",
}
