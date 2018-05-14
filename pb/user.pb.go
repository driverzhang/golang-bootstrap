// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// 请求
type UserRq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRq) Reset()         { *m = UserRq{} }
func (m *UserRq) String() string { return proto.CompactTextString(m) }
func (*UserRq) ProtoMessage()    {}
func (*UserRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_64c5967ac69a6fd2, []int{0}
}
func (m *UserRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRq.Unmarshal(m, b)
}
func (m *UserRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRq.Marshal(b, m, deterministic)
}
func (dst *UserRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRq.Merge(dst, src)
}
func (m *UserRq) XXX_Size() int {
	return xxx_messageInfo_UserRq.Size(m)
}
func (m *UserRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRq proto.InternalMessageInfo

func (m *UserRq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 响应
type UserRp struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRp) Reset()         { *m = UserRp{} }
func (m *UserRp) String() string { return proto.CompactTextString(m) }
func (*UserRp) ProtoMessage()    {}
func (*UserRp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_64c5967ac69a6fd2, []int{1}
}
func (m *UserRp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRp.Unmarshal(m, b)
}
func (m *UserRp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRp.Marshal(b, m, deterministic)
}
func (dst *UserRp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRp.Merge(dst, src)
}
func (m *UserRp) XXX_Size() int {
	return xxx_messageInfo_UserRp.Size(m)
}
func (m *UserRp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRp proto.InternalMessageInfo

func (m *UserRp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRq)(nil), "pb.UserRq")
	proto.RegisterType((*UserRp)(nil), "pb.UserRp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Data service

type DataClient interface {
	GetUser(ctx context.Context, in *UserRq, opts ...grpc.CallOption) (*UserRp, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) GetUser(ctx context.Context, in *UserRq, opts ...grpc.CallOption) (*UserRp, error) {
	out := new(UserRp)
	err := grpc.Invoke(ctx, "/pb.Data/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Data service

type DataServer interface {
	GetUser(context.Context, *UserRq) (*UserRp, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Data/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetUser(ctx, req.(*UserRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Data_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_64c5967ac69a6fd2) }

var fileDescriptor_user_64c5967ac69a6fd2 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe0, 0x62, 0x0b, 0x2d,
	0x4e, 0x2d, 0x0a, 0x2a, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0d, 0x62, 0xca, 0x4c, 0x51, 0x92, 0x81, 0xca, 0x14, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6,
	0xa6, 0x82, 0xe5, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x4d, 0x2e, 0x16, 0x97, 0xc4, 0x92, 0x44, 0x21,
	0x45, 0x2e, 0x76, 0xf7, 0xd4, 0x12, 0x90, 0x42, 0x21, 0x2e, 0xbd, 0x82, 0x24, 0x3d, 0x88, 0x61,
	0x52, 0x08, 0x76, 0x41, 0x12, 0x1b, 0xd8, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0xde, 0xe2, 0x50, 0x7b, 0x00, 0x00, 0x00,
}
