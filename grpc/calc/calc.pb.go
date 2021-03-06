// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc/calc.proto

package calc

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

type Numbers struct {
	Numbers              []int32  `protobuf:"varint,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_57803c397abb7e3e, []int{0}
}
func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (dst *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(dst, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type Result struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_57803c397abb7e3e, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "calc.Numbers")
	proto.RegisterType((*Result)(nil), "calc.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Result, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/calc.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Add(context.Context, *Numbers) (*Result, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc/calc.proto",
}

func init() { proto.RegisterFile("calc/calc.proto", fileDescriptor_calc_57803c397abb7e3e) }

var fileDescriptor_calc_57803c397abb7e3e = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0xcc, 0x49,
	0xd6, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x32, 0x17,
	0xbb, 0x5f, 0x69, 0x6e, 0x52, 0x6a, 0x51, 0xb1, 0x90, 0x04, 0x17, 0x7b, 0x1e, 0x84, 0x29, 0xc1,
	0xa8, 0xc0, 0xac, 0xc1, 0x1a, 0x04, 0xe3, 0x2a, 0x29, 0x70, 0xb1, 0x05, 0xa5, 0x16, 0x97, 0xe6,
	0x94, 0x08, 0x89, 0x71, 0xb1, 0x15, 0x81, 0x59, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50,
	0x9e, 0x91, 0x11, 0x17, 0x97, 0x73, 0x62, 0x4e, 0x72, 0x69, 0x4e, 0x62, 0x49, 0x7e, 0x91, 0x90,
	0x0a, 0x17, 0xb3, 0x63, 0x4a, 0x8a, 0x10, 0xaf, 0x1e, 0xd8, 0x3a, 0xa8, 0xf9, 0x52, 0x3c, 0x10,
	0x2e, 0xc4, 0x24, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x3b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0xba, 0x16, 0xe5, 0x9a, 0x00, 0x00, 0x00,
}
