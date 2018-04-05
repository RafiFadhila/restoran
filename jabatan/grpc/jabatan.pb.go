// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jabatan.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	jabatan.proto

It has these top-level messages:
	AddJabatanReq
	ReadJabatan
	ReadJabatanAktifResp
	ReadJabatanResp
	UpdateJabatanReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddJabatanReq struct {
	Jabatan   string `protobuf:"bytes,1,opt,name=Jabatan" json:"Jabatan,omitempty"`
	Gaji      int64  `protobuf:"varint,2,opt,name=Gaji" json:"Gaji,omitempty"`
	Status    int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,5,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,6,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdateOn  string `protobuf:"bytes,7,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
}

func (m *AddJabatanReq) Reset()                    { *m = AddJabatanReq{} }
func (m *AddJabatanReq) String() string            { return proto.CompactTextString(m) }
func (*AddJabatanReq) ProtoMessage()               {}
func (*AddJabatanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddJabatanReq) GetJabatan() string {
	if m != nil {
		return m.Jabatan
	}
	return ""
}

func (m *AddJabatanReq) GetGaji() int64 {
	if m != nil {
		return m.Gaji
	}
	return 0
}

func (m *AddJabatanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddJabatanReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddJabatanReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddJabatanReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *AddJabatanReq) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

type ReadJabatan struct {
	IDJabatan string `protobuf:"bytes,1,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
	Jabatan   string `protobuf:"bytes,2,opt,name=jabatan" json:"jabatan,omitempty"`
	Gaji      int64  `protobuf:"varint,3,opt,name=Gaji" json:"Gaji,omitempty"`
	Status    int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,5,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,6,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,7,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdateOn  string `protobuf:"bytes,8,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
}

func (m *ReadJabatan) Reset()                    { *m = ReadJabatan{} }
func (m *ReadJabatan) String() string            { return proto.CompactTextString(m) }
func (*ReadJabatan) ProtoMessage()               {}
func (*ReadJabatan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadJabatan) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

func (m *ReadJabatan) GetJabatan() string {
	if m != nil {
		return m.Jabatan
	}
	return ""
}

func (m *ReadJabatan) GetGaji() int64 {
	if m != nil {
		return m.Gaji
	}
	return 0
}

func (m *ReadJabatan) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadJabatan) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadJabatan) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *ReadJabatan) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *ReadJabatan) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

type ReadJabatanAktifResp struct {
	AllStatus []*ReadJabatan `protobuf:"bytes,1,rep,name=allStatus" json:"allStatus,omitempty"`
}

func (m *ReadJabatanAktifResp) Reset()                    { *m = ReadJabatanAktifResp{} }
func (m *ReadJabatanAktifResp) String() string            { return proto.CompactTextString(m) }
func (*ReadJabatanAktifResp) ProtoMessage()               {}
func (*ReadJabatanAktifResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadJabatanAktifResp) GetAllStatus() []*ReadJabatan {
	if m != nil {
		return m.AllStatus
	}
	return nil
}

type ReadJabatanResp struct {
	AllJabatan []*ReadJabatan `protobuf:"bytes,1,rep,name=allJabatan" json:"allJabatan,omitempty"`
}

func (m *ReadJabatanResp) Reset()                    { *m = ReadJabatanResp{} }
func (m *ReadJabatanResp) String() string            { return proto.CompactTextString(m) }
func (*ReadJabatanResp) ProtoMessage()               {}
func (*ReadJabatanResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadJabatanResp) GetAllJabatan() []*ReadJabatan {
	if m != nil {
		return m.AllJabatan
	}
	return nil
}

type UpdateJabatanReq struct {
	IDJabatan string `protobuf:"bytes,1,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
	Jabatan   string `protobuf:"bytes,2,opt,name=Jabatan" json:"Jabatan,omitempty"`
	Gaji      int64  `protobuf:"varint,3,opt,name=Gaji" json:"Gaji,omitempty"`
	Status    int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,5,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,6,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,7,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdateOn  string `protobuf:"bytes,8,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
}

func (m *UpdateJabatanReq) Reset()                    { *m = UpdateJabatanReq{} }
func (m *UpdateJabatanReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateJabatanReq) ProtoMessage()               {}
func (*UpdateJabatanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateJabatanReq) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

func (m *UpdateJabatanReq) GetJabatan() string {
	if m != nil {
		return m.Jabatan
	}
	return ""
}

func (m *UpdateJabatanReq) GetGaji() int64 {
	if m != nil {
		return m.Gaji
	}
	return 0
}

func (m *UpdateJabatanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateJabatanReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateJabatanReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateJabatanReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *UpdateJabatanReq) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

func init() {
	proto.RegisterType((*AddJabatanReq)(nil), "grpc.AddJabatanReq")
	proto.RegisterType((*ReadJabatan)(nil), "grpc.ReadJabatan")
	proto.RegisterType((*ReadJabatanAktifResp)(nil), "grpc.ReadJabatanAktifResp")
	proto.RegisterType((*ReadJabatanResp)(nil), "grpc.ReadJabatanResp")
	proto.RegisterType((*UpdateJabatanReq)(nil), "grpc.UpdateJabatanReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for JabatanService service

type JabatanServiceClient interface {
	AddJabatan(ctx context.Context, in *AddJabatanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadJabatanAktif(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadJabatanAktifResp, error)
	ReadJabatan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadJabatanResp, error)
	UpdateJabatan(ctx context.Context, in *UpdateJabatanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type jabatanServiceClient struct {
	cc *grpc1.ClientConn
}

func NewJabatanServiceClient(cc *grpc1.ClientConn) JabatanServiceClient {
	return &jabatanServiceClient{cc}
}

func (c *jabatanServiceClient) AddJabatan(ctx context.Context, in *AddJabatanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.JabatanService/AddJabatan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jabatanServiceClient) ReadJabatanAktif(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadJabatanAktifResp, error) {
	out := new(ReadJabatanAktifResp)
	err := grpc1.Invoke(ctx, "/grpc.JabatanService/ReadJabatanAktif", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jabatanServiceClient) ReadJabatan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadJabatanResp, error) {
	out := new(ReadJabatanResp)
	err := grpc1.Invoke(ctx, "/grpc.JabatanService/ReadJabatan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jabatanServiceClient) UpdateJabatan(ctx context.Context, in *UpdateJabatanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.JabatanService/UpdateJabatan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JabatanService service

type JabatanServiceServer interface {
	AddJabatan(context.Context, *AddJabatanReq) (*google_protobuf.Empty, error)
	ReadJabatanAktif(context.Context, *google_protobuf.Empty) (*ReadJabatanAktifResp, error)
	ReadJabatan(context.Context, *google_protobuf.Empty) (*ReadJabatanResp, error)
	UpdateJabatan(context.Context, *UpdateJabatanReq) (*google_protobuf.Empty, error)
}

func RegisterJabatanServiceServer(s *grpc1.Server, srv JabatanServiceServer) {
	s.RegisterService(&_JabatanService_serviceDesc, srv)
}

func _JabatanService_AddJabatan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJabatanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JabatanServiceServer).AddJabatan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.JabatanService/AddJabatan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JabatanServiceServer).AddJabatan(ctx, req.(*AddJabatanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JabatanService_ReadJabatanAktif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JabatanServiceServer).ReadJabatanAktif(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.JabatanService/ReadJabatanAktif",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JabatanServiceServer).ReadJabatanAktif(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JabatanService_ReadJabatan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JabatanServiceServer).ReadJabatan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.JabatanService/ReadJabatan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JabatanServiceServer).ReadJabatan(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JabatanService_UpdateJabatan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJabatanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JabatanServiceServer).UpdateJabatan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.JabatanService/UpdateJabatan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JabatanServiceServer).UpdateJabatan(ctx, req.(*UpdateJabatanReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _JabatanService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.JabatanService",
	HandlerType: (*JabatanServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddJabatan",
			Handler:    _JabatanService_AddJabatan_Handler,
		},
		{
			MethodName: "ReadJabatanAktif",
			Handler:    _JabatanService_ReadJabatanAktif_Handler,
		},
		{
			MethodName: "ReadJabatan",
			Handler:    _JabatanService_ReadJabatan_Handler,
		},
		{
			MethodName: "UpdateJabatan",
			Handler:    _JabatanService_UpdateJabatan_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "jabatan.proto",
}

func init() { proto.RegisterFile("jabatan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x4b, 0x6e, 0xe2, 0x40,
	0x10, 0x86, 0x69, 0x6c, 0x0c, 0x2e, 0xc4, 0x0c, 0xd3, 0x33, 0x83, 0x2c, 0xcf, 0x2c, 0x2c, 0xaf,
	0xbc, 0x32, 0x0a, 0x59, 0x46, 0x8a, 0x44, 0x42, 0x44, 0xc2, 0x06, 0xa9, 0x51, 0x0e, 0xd0, 0xe0,
	0x06, 0x41, 0x1c, 0xdb, 0x31, 0x26, 0x12, 0x57, 0xc9, 0x95, 0x72, 0x90, 0x6c, 0x72, 0x88, 0xa8,
	0xdd, 0x7e, 0x42, 0x4c, 0xd6, 0xd9, 0xb9, 0x1e, 0x5d, 0xaa, 0xaf, 0x7e, 0xff, 0xd0, 0xd9, 0xd0,
	0x39, 0x8d, 0xa8, 0x67, 0x07, 0xa1, 0x1f, 0xf9, 0x58, 0x5e, 0x85, 0xc1, 0x42, 0xff, 0xb7, 0xf2,
	0xfd, 0x95, 0xcb, 0xfa, 0x71, 0x6e, 0xbe, 0x5b, 0xf6, 0xd9, 0x63, 0x10, 0xed, 0x45, 0x8b, 0xf9,
	0x8a, 0xa0, 0x33, 0x74, 0x9c, 0x89, 0x78, 0x47, 0xd8, 0x13, 0xd6, 0xa0, 0x99, 0x44, 0x1a, 0x32,
	0x90, 0xa5, 0x92, 0x34, 0xc4, 0x18, 0xe4, 0x31, 0xdd, 0xac, 0xb5, 0xba, 0x81, 0x2c, 0x89, 0xc4,
	0xdf, 0xb8, 0x07, 0xca, 0x2c, 0xa2, 0xd1, 0x6e, 0xab, 0x49, 0x06, 0xb2, 0x1a, 0x24, 0x89, 0xf0,
	0x7f, 0x50, 0xaf, 0x43, 0x46, 0x23, 0xe6, 0x5c, 0xed, 0x35, 0x39, 0x9e, 0x93, 0x27, 0x0a, 0xd5,
	0xa9, 0xa7, 0x35, 0x4a, 0xd5, 0xa9, 0xc7, 0xab, 0xf7, 0x81, 0x93, 0xbc, 0x55, 0x44, 0x35, 0x4b,
	0x60, 0x1d, 0x5a, 0x22, 0x98, 0x7a, 0x5a, 0x33, 0x2e, 0x66, 0xb1, 0xf9, 0x86, 0xa0, 0x4d, 0x18,
	0x4d, 0x71, 0xf8, 0xa4, 0xbb, 0x51, 0x99, 0x26, 0x4f, 0x70, 0xd2, 0xe4, 0x5e, 0x31, 0x92, 0x4a,
	0xd2, 0x30, 0x23, 0x95, 0x3e, 0x25, 0x95, 0xab, 0x49, 0x1b, 0x27, 0x49, 0x95, 0x93, 0xa4, 0xcd,
	0x53, 0xa4, 0xad, 0x03, 0xd2, 0x31, 0xfc, 0x29, 0x80, 0x0e, 0x1f, 0xa2, 0xf5, 0x92, 0xb0, 0x6d,
	0x80, 0xfb, 0xa0, 0x52, 0xd7, 0x4d, 0x16, 0x45, 0x86, 0x64, 0xb5, 0x07, 0xbf, 0x6c, 0xfe, 0x1b,
	0xd8, 0x85, 0x76, 0x92, 0xf7, 0x98, 0x23, 0xf8, 0x59, 0xac, 0xf0, 0x19, 0x67, 0x00, 0xd4, 0x75,
	0xf3, 0xb3, 0x55, 0x0c, 0x29, 0x34, 0x99, 0xef, 0x08, 0xba, 0x62, 0xb7, 0xc2, 0x9f, 0xf4, 0xe5,
	0xf5, 0x27, 0xe5, 0xeb, 0x4f, 0xbe, 0xcb, 0xf5, 0x07, 0x2f, 0x75, 0xf8, 0x91, 0x6c, 0x3b, 0x63,
	0xe1, 0xf3, 0x7a, 0xc1, 0xf0, 0x05, 0x40, 0xee, 0x23, 0xfc, 0x5b, 0x9c, 0xab, 0xe4, 0x2c, 0xbd,
	0x67, 0x0b, 0x27, 0xda, 0xa9, 0x13, 0xed, 0x1b, 0xee, 0x44, 0xb3, 0x86, 0x6f, 0xa1, 0x7b, 0xa8,
	0x26, 0xae, 0xe8, 0xd6, 0xf5, 0x23, 0x25, 0x32, 0xf5, 0xcd, 0x1a, 0xbe, 0x2c, 0x1b, 0xa0, 0x6a,
	0xc8, 0xdf, 0x63, 0x39, 0xc5, 0xfb, 0x21, 0x74, 0x4a, 0x3a, 0xf2, 0x09, 0xbc, 0xf3, 0x50, 0xdc,
	0x6a, 0x98, 0xb9, 0x12, 0x67, 0xce, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xc8, 0xaa, 0xf5,
	0x8d, 0x04, 0x00, 0x00,
}
