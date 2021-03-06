// Code generated by protoc-gen-go. DO NOT EDIT.
// source: happy.proto

/*
Package happy is a generated protocol buffer package.

It is generated from these files:
	happy.proto

It has these top-level messages:
	ListTagRequest
	ListTagResponse
	GetRequest
	GetResponse
	PostRequest
	PostResponse
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

type ListTagRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Page  string `protobuf:"bytes,2,opt,name=page" json:"page,omitempty"`
	Num   string `protobuf:"bytes,3,opt,name=num" json:"num,omitempty"`
}

func (m *ListTagRequest) Reset()                    { *m = ListTagRequest{} }
func (m *ListTagRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTagRequest) ProtoMessage()               {}
func (*ListTagRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListTagRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ListTagRequest) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *ListTagRequest) GetNum() string {
	if m != nil {
		return m.Num
	}
	return ""
}

type ListTagResponse struct {
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Num int32  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *ListTagResponse) Reset()                    { *m = ListTagResponse{} }
func (m *ListTagResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTagResponse) ProtoMessage()               {}
func (*ListTagResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListTagResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ListTagResponse) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GetRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Tag   string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type GetResponse struct {
	Word string `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type PostRequest struct {
	Key  string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Word string `protobuf:"bytes,2,opt,name=word" json:"word,omitempty"`
	Tag  string `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PostRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PostRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *PostRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type PostResponse struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *PostResponse) Reset()                    { *m = PostResponse{} }
func (m *PostResponse) String() string            { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()               {}
func (*PostResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PostResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*ListTagRequest)(nil), "ListTagRequest")
	proto.RegisterType((*ListTagResponse)(nil), "ListTagResponse")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*PostRequest)(nil), "PostRequest")
	proto.RegisterType((*PostResponse)(nil), "PostResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HappyService service

type HappyServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (HappyService_ListTagClient, error)
}

type happyServiceClient struct {
	cc *grpc.ClientConn
}

func NewHappyServiceClient(cc *grpc.ClientConn) HappyServiceClient {
	return &happyServiceClient{cc}
}

func (c *happyServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/HappyService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *happyServiceClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := grpc.Invoke(ctx, "/HappyService/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *happyServiceClient) ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (HappyService_ListTagClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HappyService_serviceDesc.Streams[0], c.cc, "/HappyService/ListTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &happyServiceListTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HappyService_ListTagClient interface {
	Recv() (*ListTagResponse, error)
	grpc.ClientStream
}

type happyServiceListTagClient struct {
	grpc.ClientStream
}

func (x *happyServiceListTagClient) Recv() (*ListTagResponse, error) {
	m := new(ListTagResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HappyService service

type HappyServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	ListTag(*ListTagRequest, HappyService_ListTagServer) error
}

func RegisterHappyServiceServer(s *grpc.Server, srv HappyServiceServer) {
	s.RegisterService(&_HappyService_serviceDesc, srv)
}

func _HappyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
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
		return srv.(HappyServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HappyService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HappyServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HappyService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HappyServiceServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HappyService_ListTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTagRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HappyServiceServer).ListTag(m, &happyServiceListTagServer{stream})
}

type HappyService_ListTagServer interface {
	Send(*ListTagResponse) error
	grpc.ServerStream
}

type happyServiceListTagServer struct {
	grpc.ServerStream
}

func (x *happyServiceListTagServer) Send(m *ListTagResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HappyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HappyService",
	HandlerType: (*HappyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _HappyService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _HappyService_Post_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTag",
			Handler:       _HappyService_ListTag_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "happy.proto",
}

func init() { proto.RegisterFile("happy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x49, 0xd2, 0xfe, 0x7e, 0x38, 0x89, 0x6d, 0x18, 0x8a, 0xd4, 0x20, 0x58, 0xf7, 0x24,
	0x1e, 0x36, 0xfe, 0x05, 0xf1, 0xac, 0x54, 0xa4, 0x07, 0xa9, 0xbe, 0xc0, 0xaa, 0x43, 0x0c, 0xb5,
	0xd9, 0x34, 0xd9, 0x56, 0x8a, 0x78, 0xf1, 0x15, 0x7c, 0x34, 0x6f, 0x9e, 0x7d, 0x10, 0xd9, 0x4d,
	0x9a, 0xb4, 0x3d, 0x78, 0xfb, 0x66, 0x93, 0xcf, 0x67, 0x66, 0x32, 0x0b, 0xee, 0xb3, 0x48, 0xd3,
	0x39, 0x4f, 0x33, 0xa9, 0x64, 0xb0, 0x13, 0x49, 0x19, 0xbd, 0x50, 0x28, 0xd2, 0x38, 0x14, 0x49,
	0x22, 0x95, 0x50, 0xb1, 0x4c, 0xf2, 0xe2, 0x2d, 0x1b, 0x40, 0x6b, 0x10, 0xe7, 0xea, 0x5e, 0x44,
	0x43, 0x9a, 0x4c, 0x29, 0x57, 0xd8, 0x81, 0xe6, 0x64, 0x4a, 0xd9, 0xbc, 0x6b, 0xf5, 0xac, 0xfd,
	0x8d, 0x61, 0xf1, 0x80, 0x08, 0x8d, 0x54, 0x44, 0xd4, 0xb5, 0xcd, 0xa1, 0xc9, 0xe8, 0x83, 0x93,
	0x4c, 0xc7, 0x5d, 0xc7, 0x1c, 0xe9, 0xc8, 0xce, 0xa0, 0x5d, 0xd9, 0xf2, 0x54, 0x26, 0xb9, 0xf9,
	0x48, 0x89, 0xa8, 0x94, 0xe9, 0xb8, 0xc0, 0xb4, 0xa9, 0x59, 0x60, 0xa7, 0x00, 0x7d, 0x52, 0x7f,
	0x37, 0x50, 0x7a, 0xec, 0xca, 0xc3, 0xf6, 0xc0, 0x35, 0x54, 0x59, 0x08, 0xa1, 0xf1, 0x2a, 0xb3,
	0xa7, 0x92, 0x32, 0x99, 0x5d, 0x81, 0x7b, 0x2b, 0xf3, 0xca, 0xec, 0x83, 0x33, 0xa2, 0x85, 0x57,
	0xc7, 0x0a, 0xb2, 0x6b, 0x68, 0x51, 0xc9, 0xa9, 0x2b, 0xf5, 0xc0, 0x2b, 0x34, 0xf5, 0x4c, 0xab,
	0x9e, 0xe3, 0x6f, 0x0b, 0xbc, 0x6b, 0xfd, 0xd3, 0xef, 0x28, 0x9b, 0xc5, 0x8f, 0x84, 0xe7, 0xe0,
	0xf4, 0x49, 0xa1, 0xcb, 0xeb, 0xc1, 0x02, 0x8f, 0x2f, 0xf5, 0xcb, 0xb6, 0x3e, 0xbe, 0x7e, 0x3e,
	0x6d, 0x1f, 0x5b, 0xe1, 0xec, 0x28, 0xa4, 0xb1, 0x54, 0xa1, 0xd9, 0x1a, 0x5e, 0x42, 0x43, 0x17,
	0x43, 0x8f, 0x2f, 0xb5, 0x1e, 0x6c, 0xf2, 0xe5, 0x0e, 0xd8, 0xae, 0x81, 0xb7, 0x59, 0x67, 0x15,
	0x0e, 0xdf, 0x46, 0x34, 0x7f, 0xbf, 0xb0, 0x0e, 0xf0, 0x06, 0xfe, 0x97, 0x9b, 0xc0, 0x36, 0x5f,
	0xdd, 0x70, 0xe0, 0xf3, 0xb5, 0x25, 0xb1, 0xc0, 0xe8, 0x3a, 0x88, 0x6b, 0x3a, 0x25, 0xa2, 0x43,
	0xeb, 0xe1, 0x9f, 0xb9, 0x2a, 0x27, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x5c, 0xd6, 0xa7,
	0x57, 0x02, 0x00, 0x00,
}
