// Code generated by protoc-gen-go. DO NOT EDIT.
// source: happy.proto

/*
Package happy is a generated protocol buffer package.

It is generated from these files:
	happy.proto

It has these top-level messages:
	HappyRequest
	HappyResponse
*/
package happy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HappyRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *HappyRequest) Reset()                    { *m = HappyRequest{} }
func (m *HappyRequest) String() string            { return proto.CompactTextString(m) }
func (*HappyRequest) ProtoMessage()               {}
func (*HappyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HappyRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type HappyResponse struct {
	Word string `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
}

func (m *HappyResponse) Reset()                    { *m = HappyResponse{} }
func (m *HappyResponse) String() string            { return proto.CompactTextString(m) }
func (*HappyResponse) ProtoMessage()               {}
func (*HappyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HappyResponse) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func init() {
	proto.RegisterType((*HappyRequest)(nil), "HappyRequest")
	proto.RegisterType((*HappyResponse)(nil), "HappyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HappyService service

type HappyServiceClient interface {
	Get(ctx context.Context, in *HappyRequest, opts ...grpc.CallOption) (*HappyResponse, error)
}

type happyServiceClient struct {
	cc *grpc.ClientConn
}

func NewHappyServiceClient(cc *grpc.ClientConn) HappyServiceClient {
	return &happyServiceClient{cc}
}

func (c *happyServiceClient) Get(ctx context.Context, in *HappyRequest, opts ...grpc.CallOption) (*HappyResponse, error) {
	out := new(HappyResponse)
	err := grpc.Invoke(ctx, "/HappyService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HappyService service

type HappyServiceServer interface {
	Get(context.Context, *HappyRequest) (*HappyResponse, error)
}

func RegisterHappyServiceServer(s *grpc.Server, srv HappyServiceServer) {
	s.RegisterService(&_HappyService_serviceDesc, srv)
}

func _HappyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HappyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HappyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HappyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HappyServiceServer).Get(ctx, req.(*HappyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HappyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HappyService",
	HandlerType: (*HappyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _HappyService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "happy.proto",
}

func init() { proto.RegisterFile("happy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0x2c, 0x28,
	0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86,
	0xc8, 0x2a, 0xa9, 0x70, 0xf1, 0x78, 0x80, 0x14, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x89, 0x70, 0xb1, 0x16, 0x96, 0xa6, 0x16, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41,
	0x38, 0x4a, 0xca, 0x5c, 0xbc, 0x50, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c,
	0x2c, 0xe5, 0xf9, 0x45, 0x29, 0x50, 0x55, 0x60, 0xb6, 0x91, 0x0f, 0xd4, 0xa8, 0xe0, 0xd4, 0xa2,
	0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x1b, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x5e, 0x3d, 0x64, 0x0b,
	0xa4, 0xf8, 0xf4, 0x50, 0x4c, 0x52, 0x12, 0x6b, 0xba, 0xfc, 0x64, 0x32, 0x93, 0x80, 0x10, 0x9f,
	0x7e, 0x99, 0xa1, 0x7e, 0x6a, 0x6e, 0x7e, 0x89, 0x3e, 0xd8, 0xf1, 0x49, 0x6c, 0x60, 0xf7, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x89, 0xf5, 0xa0, 0xcc, 0x00, 0x00, 0x00,
}
