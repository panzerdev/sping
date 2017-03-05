// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Time
	Ping
	Pong
*/
package echo

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

type Time struct {
	Seconds     int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Time) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Time) GetNanoseconds() int64 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type Ping struct {
	Sender string `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Sseq   int64  `protobuf:"varint,2,opt,name=sseq" json:"sseq,omitempty"`
	Sent   *Time  `protobuf:"bytes,3,opt,name=sent" json:"sent,omitempty"`
	Ttl    int64  `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ping) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Ping) GetSseq() int64 {
	if m != nil {
		return m.Sseq
	}
	return 0
}

func (m *Ping) GetSent() *Time {
	if m != nil {
		return m.Sent
	}
	return nil
}

func (m *Ping) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type Pong struct {
	Success bool  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Sseq    int64 `protobuf:"varint,2,opt,name=sseq" json:"sseq,omitempty"`
	Rseq    int64 `protobuf:"varint,3,opt,name=rseq" json:"rseq,omitempty"`
	Sent    *Time `protobuf:"bytes,4,opt,name=sent" json:"sent,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Pong) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Pong) GetSseq() int64 {
	if m != nil {
		return m.Sseq
	}
	return 0
}

func (m *Pong) GetRseq() int64 {
	if m != nil {
		return m.Rseq
	}
	return 0
}

func (m *Pong) GetSent() *Time {
	if m != nil {
		return m.Sent
	}
	return nil
}

func init() {
	proto.RegisterType((*Time)(nil), "echo.Time")
	proto.RegisterType((*Ping)(nil), "echo.Ping")
	proto.RegisterType((*Pong)(nil), "echo.Pong")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SecurePing service

type SecurePingClient interface {
	Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type securePingClient struct {
	cc *grpc.ClientConn
}

func NewSecurePingClient(cc *grpc.ClientConn) SecurePingClient {
	return &securePingClient{cc}
}

func (c *securePingClient) Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/echo.SecurePing/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecurePing service

type SecurePingServer interface {
	Echo(context.Context, *Ping) (*Pong, error)
}

func RegisterSecurePingServer(s *grpc.Server, srv SecurePingServer) {
	s.RegisterService(&_SecurePing_serviceDesc, srv)
}

func _SecurePing_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurePingServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.SecurePing/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurePingServer).Echo(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurePing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.SecurePing",
	HandlerType: (*SecurePingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _SecurePing_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x40,
	0x10, 0x44, 0x31, 0x5e, 0x05, 0x98, 0x34, 0x68, 0x0b, 0x64, 0x51, 0x20, 0xcb, 0x55, 0x2a, 0x17,
	0xe1, 0x0f, 0x90, 0xe8, 0x23, 0xc3, 0x0f, 0xc0, 0x79, 0x65, 0x47, 0x82, 0x5d, 0xb8, 0x73, 0xfe,
	0x3f, 0xba, 0x8d, 0x2d, 0xa5, 0x70, 0x37, 0xf3, 0xf6, 0xf6, 0x66, 0xee, 0x00, 0x09, 0xa3, 0xb5,
	0x7f, 0xd1, 0x26, 0x63, 0xca, 0xba, 0x79, 0x03, 0x7d, 0x1e, 0x7f, 0x85, 0x2b, 0xdc, 0x25, 0x09,
	0xa6, 0x7d, 0xaa, 0x8a, 0xba, 0xd8, 0x95, 0xdd, 0x62, 0xb9, 0xc6, 0x56, 0xbf, 0xd4, 0x96, 0xe9,
	0xad, 0x4f, 0xaf, 0x51, 0xd3, 0x83, 0x0e, 0x47, 0x1d, 0xf8, 0x09, 0x9b, 0x24, 0xda, 0x4b, 0xf4,
	0x2b, 0x1e, 0xba, 0xd9, 0x31, 0x83, 0x52, 0x92, 0xff, 0x79, 0xd5, 0x35, 0xbf, 0x80, 0x92, 0xe8,
	0x54, 0x95, 0x75, 0xb1, 0xdb, 0xee, 0xd1, 0x7a, 0xb1, 0xdc, 0xa4, 0x73, 0xce, 0x8f, 0x28, 0xa7,
	0xe9, 0xa7, 0x22, 0x5f, 0xc9, 0xb2, 0x19, 0x41, 0x07, 0xd3, 0xc1, 0x9b, 0x9e, 0x42, 0x90, 0x74,
	0x69, 0x7a, 0xdf, 0x2d, 0x76, 0x35, 0x87, 0x41, 0x31, 0xb3, 0xf2, 0xc2, 0xe2, 0x75, 0x36, 0xad,
	0x67, 0xef, 0x5b, 0xe0, 0x43, 0xc2, 0x29, 0x8a, 0xbf, 0xaa, 0x06, 0xbd, 0x87, 0xd1, 0x78, 0x3e,
	0x97, 0xd9, 0xf3, 0xa2, 0x4d, 0x87, 0xe6, 0xe6, 0x7b, 0xe3, 0x1f, 0xfa, 0x7a, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0xd1, 0x21, 0x44, 0x5e, 0x01, 0x00, 0x00,
}