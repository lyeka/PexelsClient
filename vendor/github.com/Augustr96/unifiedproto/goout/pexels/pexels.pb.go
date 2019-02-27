// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/pexels/pexels.proto

package pexels

import (
	context "context"
	fmt "fmt"
	utils "github.com/Augustr96/unifiedproto/goout/utils"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetPhotoRequest struct {
	// photo id
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPhotoRequest) Reset()         { *m = GetPhotoRequest{} }
func (m *GetPhotoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPhotoRequest) ProtoMessage()    {}
func (*GetPhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0204bdbded08973, []int{0}
}

func (m *GetPhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhotoRequest.Unmarshal(m, b)
}
func (m *GetPhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhotoRequest.Marshal(b, m, deterministic)
}
func (m *GetPhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhotoRequest.Merge(m, src)
}
func (m *GetPhotoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPhotoRequest.Size(m)
}
func (m *GetPhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhotoRequest proto.InternalMessageInfo

func (m *GetPhotoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPhotoResponse struct {
	Status               *utils.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPhotoResponse) Reset()         { *m = GetPhotoResponse{} }
func (m *GetPhotoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPhotoResponse) ProtoMessage()    {}
func (*GetPhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0204bdbded08973, []int{1}
}

func (m *GetPhotoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhotoResponse.Unmarshal(m, b)
}
func (m *GetPhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhotoResponse.Marshal(b, m, deterministic)
}
func (m *GetPhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhotoResponse.Merge(m, src)
}
func (m *GetPhotoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPhotoResponse.Size(m)
}
func (m *GetPhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhotoResponse proto.InternalMessageInfo

func (m *GetPhotoResponse) GetStatus() *utils.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPhotoRequest)(nil), "pexels.GetPhotoRequest")
	proto.RegisterType((*GetPhotoResponse)(nil), "pexels.GetPhotoResponse")
}

func init() { proto.RegisterFile("src/pexels/pexels.proto", fileDescriptor_e0204bdbded08973) }

var fileDescriptor_e0204bdbded08973 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2e, 0x4a, 0xd6,
	0x2f, 0x48, 0xad, 0x48, 0xcd, 0x29, 0x86, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c,
	0x10, 0x9e, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62,
	0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x54, 0x95, 0x94, 0x28, 0x48, 0x7b, 0x69,
	0x49, 0x66, 0x4e, 0x31, 0x84, 0x84, 0x08, 0x2b, 0x29, 0x72, 0xf1, 0xbb, 0xa7, 0x96, 0x04, 0x64,
	0xe4, 0x97, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x59, 0x72, 0x09, 0x20, 0x94,
	0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xa9, 0x72, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16,
	0x83, 0xd5, 0x71, 0x1b, 0xf1, 0xea, 0x41, 0x0c, 0x0d, 0x06, 0x0b, 0x06, 0x41, 0x25, 0x8d, 0x12,
	0xb9, 0xd8, 0x02, 0xc0, 0x8e, 0x13, 0x0a, 0xe7, 0xe2, 0x80, 0x19, 0x22, 0x24, 0xae, 0x07, 0x75,
	0x3f, 0x9a, 0xcd, 0x52, 0x12, 0x98, 0x12, 0x10, 0xfb, 0x94, 0x24, 0x9a, 0x2e, 0x3f, 0x99, 0xcc,
	0x24, 0x24, 0x24, 0x00, 0x0b, 0x01, 0x98, 0x0a, 0x27, 0x83, 0x28, 0xbd, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xd2, 0xf4, 0xd2, 0xe2, 0x92, 0x22, 0x4b, 0x33,
	0xfd, 0xd2, 0xbc, 0xcc, 0xb4, 0xcc, 0xd4, 0x14, 0xb0, 0x2f, 0xf5, 0xd3, 0xf3, 0xf3, 0x4b, 0x4b,
	0xa0, 0x5a, 0x93, 0xd8, 0xc0, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x65, 0xdb,
	0xdd, 0x51, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PexelsClient is the client API for Pexels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PexelsClient interface {
	GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*GetPhotoResponse, error)
}

type pexelsClient struct {
	cc *grpc.ClientConn
}

func NewPexelsClient(cc *grpc.ClientConn) PexelsClient {
	return &pexelsClient{cc}
}

func (c *pexelsClient) GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*GetPhotoResponse, error) {
	out := new(GetPhotoResponse)
	err := c.cc.Invoke(ctx, "/pexels.Pexels/GetPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PexelsServer is the server API for Pexels service.
type PexelsServer interface {
	GetPhoto(context.Context, *GetPhotoRequest) (*GetPhotoResponse, error)
}

func RegisterPexelsServer(s *grpc.Server, srv PexelsServer) {
	s.RegisterService(&_Pexels_serviceDesc, srv)
}

func _Pexels_GetPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PexelsServer).GetPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pexels.Pexels/GetPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PexelsServer).GetPhoto(ctx, req.(*GetPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pexels_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pexels.Pexels",
	HandlerType: (*PexelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhoto",
			Handler:    _Pexels_GetPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/pexels/pexels.proto",
}