// Code generated by protoc-gen-go. DO NOT EDIT.
// source: karyawan.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	karyawan.proto

It has these top-level messages:
	AddKaryawanReq
	ReadKaryawan
	ReadKaryawanByJabatanReq
	ReadKaryawanByJabatanResp
	ReadKaryawanByBagianReq
	ReadKaryawanByBagianResp
	ReadKaryawanByKeteranganReq
	ReadKaryawanByKeteranganResp
	ReadKaryawanResp
	UpdateKaryawanReq
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

type AddKaryawanReq struct {
	Nama       string `protobuf:"bytes,1,opt,name=Nama" json:"Nama,omitempty"`
	Alamat     string `protobuf:"bytes,2,opt,name=Alamat" json:"Alamat,omitempty"`
	Telepon    string `protobuf:"bytes,3,opt,name=Telepon" json:"Telepon,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	IDJabatan  string `protobuf:"bytes,5,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
	IDBagian   string `protobuf:"bytes,6,opt,name=IDBagian" json:"IDBagian,omitempty"`
	Status     int32  `protobuf:"varint,7,opt,name=Status" json:"Status,omitempty"`
	CreatedBy  string `protobuf:"bytes,8,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn  string `protobuf:"bytes,9,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy  string `protobuf:"bytes,10,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdatedOn  string `protobuf:"bytes,11,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	Keterangan string `protobuf:"bytes,12,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *AddKaryawanReq) Reset()                    { *m = AddKaryawanReq{} }
func (m *AddKaryawanReq) String() string            { return proto.CompactTextString(m) }
func (*AddKaryawanReq) ProtoMessage()               {}
func (*AddKaryawanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddKaryawanReq) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *AddKaryawanReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *AddKaryawanReq) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *AddKaryawanReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddKaryawanReq) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

func (m *AddKaryawanReq) GetIDBagian() string {
	if m != nil {
		return m.IDBagian
	}
	return ""
}

func (m *AddKaryawanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddKaryawanReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddKaryawanReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddKaryawanReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *AddKaryawanReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *AddKaryawanReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadKaryawan struct {
	IDKaryawan string `protobuf:"bytes,1,opt,name=IDKaryawan" json:"IDKaryawan,omitempty"`
	Nama       string `protobuf:"bytes,2,opt,name=Nama" json:"Nama,omitempty"`
	Alamat     string `protobuf:"bytes,3,opt,name=Alamat" json:"Alamat,omitempty"`
	Telepon    string `protobuf:"bytes,4,opt,name=Telepon" json:"Telepon,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=Email" json:"Email,omitempty"`
	IDJabatan  string `protobuf:"bytes,6,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
	IDBagian   string `protobuf:"bytes,7,opt,name=IDBagian" json:"IDBagian,omitempty"`
	Status     int32  `protobuf:"varint,8,opt,name=Status" json:"Status,omitempty"`
	CreatedBy  string `protobuf:"bytes,9,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn  string `protobuf:"bytes,10,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy  string `protobuf:"bytes,11,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdatedOn  string `protobuf:"bytes,12,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	Keterangan string `protobuf:"bytes,13,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *ReadKaryawan) Reset()                    { *m = ReadKaryawan{} }
func (m *ReadKaryawan) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawan) ProtoMessage()               {}
func (*ReadKaryawan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadKaryawan) GetIDKaryawan() string {
	if m != nil {
		return m.IDKaryawan
	}
	return ""
}

func (m *ReadKaryawan) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *ReadKaryawan) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *ReadKaryawan) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *ReadKaryawan) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReadKaryawan) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

func (m *ReadKaryawan) GetIDBagian() string {
	if m != nil {
		return m.IDBagian
	}
	return ""
}

func (m *ReadKaryawan) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadKaryawan) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadKaryawan) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *ReadKaryawan) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *ReadKaryawan) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *ReadKaryawan) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadKaryawanByJabatanReq struct {
	IDJabatan string `protobuf:"bytes,1,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
}

func (m *ReadKaryawanByJabatanReq) Reset()                    { *m = ReadKaryawanByJabatanReq{} }
func (m *ReadKaryawanByJabatanReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByJabatanReq) ProtoMessage()               {}
func (*ReadKaryawanByJabatanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadKaryawanByJabatanReq) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

type ReadKaryawanByJabatanResp struct {
	AllJabatan []*ReadKaryawan `protobuf:"bytes,1,rep,name=allJabatan" json:"allJabatan,omitempty"`
}

func (m *ReadKaryawanByJabatanResp) Reset()                    { *m = ReadKaryawanByJabatanResp{} }
func (m *ReadKaryawanByJabatanResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByJabatanResp) ProtoMessage()               {}
func (*ReadKaryawanByJabatanResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadKaryawanByJabatanResp) GetAllJabatan() []*ReadKaryawan {
	if m != nil {
		return m.AllJabatan
	}
	return nil
}

type ReadKaryawanByBagianReq struct {
	IDBagian string `protobuf:"bytes,1,opt,name=IDBagian" json:"IDBagian,omitempty"`
}

func (m *ReadKaryawanByBagianReq) Reset()                    { *m = ReadKaryawanByBagianReq{} }
func (m *ReadKaryawanByBagianReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByBagianReq) ProtoMessage()               {}
func (*ReadKaryawanByBagianReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadKaryawanByBagianReq) GetIDBagian() string {
	if m != nil {
		return m.IDBagian
	}
	return ""
}

type ReadKaryawanByBagianResp struct {
	AllBagian []*ReadKaryawan `protobuf:"bytes,1,rep,name=allBagian" json:"allBagian,omitempty"`
}

func (m *ReadKaryawanByBagianResp) Reset()                    { *m = ReadKaryawanByBagianResp{} }
func (m *ReadKaryawanByBagianResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByBagianResp) ProtoMessage()               {}
func (*ReadKaryawanByBagianResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadKaryawanByBagianResp) GetAllBagian() []*ReadKaryawan {
	if m != nil {
		return m.AllBagian
	}
	return nil
}

type ReadKaryawanByKeteranganReq struct {
	Keterangan string `protobuf:"bytes,1,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *ReadKaryawanByKeteranganReq) Reset()                    { *m = ReadKaryawanByKeteranganReq{} }
func (m *ReadKaryawanByKeteranganReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByKeteranganReq) ProtoMessage()               {}
func (*ReadKaryawanByKeteranganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadKaryawanByKeteranganReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadKaryawanByKeteranganResp struct {
	AllKeterangan []*ReadKaryawan `protobuf:"bytes,1,rep,name=allKeterangan" json:"allKeterangan,omitempty"`
}

func (m *ReadKaryawanByKeteranganResp) Reset()                    { *m = ReadKaryawanByKeteranganResp{} }
func (m *ReadKaryawanByKeteranganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanByKeteranganResp) ProtoMessage()               {}
func (*ReadKaryawanByKeteranganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadKaryawanByKeteranganResp) GetAllKeterangan() []*ReadKaryawan {
	if m != nil {
		return m.AllKeterangan
	}
	return nil
}

type ReadKaryawanResp struct {
	AllKaryawan []*ReadKaryawan `protobuf:"bytes,1,rep,name=allKaryawan" json:"allKaryawan,omitempty"`
}

func (m *ReadKaryawanResp) Reset()                    { *m = ReadKaryawanResp{} }
func (m *ReadKaryawanResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKaryawanResp) ProtoMessage()               {}
func (*ReadKaryawanResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadKaryawanResp) GetAllKaryawan() []*ReadKaryawan {
	if m != nil {
		return m.AllKaryawan
	}
	return nil
}

type UpdateKaryawanReq struct {
	IDKaryawan string `protobuf:"bytes,1,opt,name=IDKaryawan" json:"IDKaryawan,omitempty"`
	Nama       string `protobuf:"bytes,2,opt,name=Nama" json:"Nama,omitempty"`
	Alamat     string `protobuf:"bytes,3,opt,name=Alamat" json:"Alamat,omitempty"`
	Telepon    string `protobuf:"bytes,4,opt,name=Telepon" json:"Telepon,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=Email" json:"Email,omitempty"`
	IDJabatan  string `protobuf:"bytes,6,opt,name=IDJabatan" json:"IDJabatan,omitempty"`
	IDBagian   string `protobuf:"bytes,7,opt,name=IDBagian" json:"IDBagian,omitempty"`
	Status     int32  `protobuf:"varint,8,opt,name=Status" json:"Status,omitempty"`
	CreatedBy  string `protobuf:"bytes,9,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn  string `protobuf:"bytes,10,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedBy  string `protobuf:"bytes,11,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
	UpdatedOn  string `protobuf:"bytes,12,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	Keterangan string `protobuf:"bytes,13,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *UpdateKaryawanReq) Reset()                    { *m = UpdateKaryawanReq{} }
func (m *UpdateKaryawanReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateKaryawanReq) ProtoMessage()               {}
func (*UpdateKaryawanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateKaryawanReq) GetIDKaryawan() string {
	if m != nil {
		return m.IDKaryawan
	}
	return ""
}

func (m *UpdateKaryawanReq) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *UpdateKaryawanReq) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *UpdateKaryawanReq) GetTelepon() string {
	if m != nil {
		return m.Telepon
	}
	return ""
}

func (m *UpdateKaryawanReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateKaryawanReq) GetIDJabatan() string {
	if m != nil {
		return m.IDJabatan
	}
	return ""
}

func (m *UpdateKaryawanReq) GetIDBagian() string {
	if m != nil {
		return m.IDBagian
	}
	return ""
}

func (m *UpdateKaryawanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateKaryawanReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateKaryawanReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateKaryawanReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *UpdateKaryawanReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *UpdateKaryawanReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func init() {
	proto.RegisterType((*AddKaryawanReq)(nil), "grpc.AddKaryawanReq")
	proto.RegisterType((*ReadKaryawan)(nil), "grpc.ReadKaryawan")
	proto.RegisterType((*ReadKaryawanByJabatanReq)(nil), "grpc.ReadKaryawanByJabatanReq")
	proto.RegisterType((*ReadKaryawanByJabatanResp)(nil), "grpc.ReadKaryawanByJabatanResp")
	proto.RegisterType((*ReadKaryawanByBagianReq)(nil), "grpc.ReadKaryawanByBagianReq")
	proto.RegisterType((*ReadKaryawanByBagianResp)(nil), "grpc.ReadKaryawanByBagianResp")
	proto.RegisterType((*ReadKaryawanByKeteranganReq)(nil), "grpc.ReadKaryawanByKeteranganReq")
	proto.RegisterType((*ReadKaryawanByKeteranganResp)(nil), "grpc.ReadKaryawanByKeteranganResp")
	proto.RegisterType((*ReadKaryawanResp)(nil), "grpc.ReadKaryawanResp")
	proto.RegisterType((*UpdateKaryawanReq)(nil), "grpc.UpdateKaryawanReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for KaryawanService service

type KaryawanServiceClient interface {
	AddKaryawan(ctx context.Context, in *AddKaryawanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKaryawanByJabatan(ctx context.Context, in *ReadKaryawanByJabatanReq, opts ...grpc1.CallOption) (*ReadKaryawanByJabatanResp, error)
	ReadKaryawanByBagian(ctx context.Context, in *ReadKaryawanByBagianReq, opts ...grpc1.CallOption) (*ReadKaryawanByBagianResp, error)
	ReadKaryawanByKeterangan(ctx context.Context, in *ReadKaryawanByKeteranganReq, opts ...grpc1.CallOption) (*ReadKaryawanByKeteranganResp, error)
	ReadKaryawan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKaryawanResp, error)
	UpdateKaryawan(ctx context.Context, in *UpdateKaryawanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type karyawanServiceClient struct {
	cc *grpc1.ClientConn
}

func NewKaryawanServiceClient(cc *grpc1.ClientConn) KaryawanServiceClient {
	return &karyawanServiceClient{cc}
}

func (c *karyawanServiceClient) AddKaryawan(ctx context.Context, in *AddKaryawanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/AddKaryawan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karyawanServiceClient) ReadKaryawanByJabatan(ctx context.Context, in *ReadKaryawanByJabatanReq, opts ...grpc1.CallOption) (*ReadKaryawanByJabatanResp, error) {
	out := new(ReadKaryawanByJabatanResp)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/ReadKaryawanByJabatan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karyawanServiceClient) ReadKaryawanByBagian(ctx context.Context, in *ReadKaryawanByBagianReq, opts ...grpc1.CallOption) (*ReadKaryawanByBagianResp, error) {
	out := new(ReadKaryawanByBagianResp)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/ReadKaryawanByBagian", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karyawanServiceClient) ReadKaryawanByKeterangan(ctx context.Context, in *ReadKaryawanByKeteranganReq, opts ...grpc1.CallOption) (*ReadKaryawanByKeteranganResp, error) {
	out := new(ReadKaryawanByKeteranganResp)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/ReadKaryawanByKeterangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karyawanServiceClient) ReadKaryawan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKaryawanResp, error) {
	out := new(ReadKaryawanResp)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/ReadKaryawan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karyawanServiceClient) UpdateKaryawan(ctx context.Context, in *UpdateKaryawanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KaryawanService/UpdateKaryawan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KaryawanService service

type KaryawanServiceServer interface {
	AddKaryawan(context.Context, *AddKaryawanReq) (*google_protobuf.Empty, error)
	ReadKaryawanByJabatan(context.Context, *ReadKaryawanByJabatanReq) (*ReadKaryawanByJabatanResp, error)
	ReadKaryawanByBagian(context.Context, *ReadKaryawanByBagianReq) (*ReadKaryawanByBagianResp, error)
	ReadKaryawanByKeterangan(context.Context, *ReadKaryawanByKeteranganReq) (*ReadKaryawanByKeteranganResp, error)
	ReadKaryawan(context.Context, *google_protobuf.Empty) (*ReadKaryawanResp, error)
	UpdateKaryawan(context.Context, *UpdateKaryawanReq) (*google_protobuf.Empty, error)
}

func RegisterKaryawanServiceServer(s *grpc1.Server, srv KaryawanServiceServer) {
	s.RegisterService(&_KaryawanService_serviceDesc, srv)
}

func _KaryawanService_AddKaryawan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKaryawanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).AddKaryawan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/AddKaryawan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).AddKaryawan(ctx, req.(*AddKaryawanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaryawanService_ReadKaryawanByJabatan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKaryawanByJabatanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).ReadKaryawanByJabatan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/ReadKaryawanByJabatan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).ReadKaryawanByJabatan(ctx, req.(*ReadKaryawanByJabatanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaryawanService_ReadKaryawanByBagian_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKaryawanByBagianReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).ReadKaryawanByBagian(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/ReadKaryawanByBagian",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).ReadKaryawanByBagian(ctx, req.(*ReadKaryawanByBagianReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaryawanService_ReadKaryawanByKeterangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKaryawanByKeteranganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).ReadKaryawanByKeterangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/ReadKaryawanByKeterangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).ReadKaryawanByKeterangan(ctx, req.(*ReadKaryawanByKeteranganReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaryawanService_ReadKaryawan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).ReadKaryawan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/ReadKaryawan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).ReadKaryawan(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KaryawanService_UpdateKaryawan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKaryawanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaryawanServiceServer).UpdateKaryawan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KaryawanService/UpdateKaryawan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaryawanServiceServer).UpdateKaryawan(ctx, req.(*UpdateKaryawanReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KaryawanService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.KaryawanService",
	HandlerType: (*KaryawanServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddKaryawan",
			Handler:    _KaryawanService_AddKaryawan_Handler,
		},
		{
			MethodName: "ReadKaryawanByJabatan",
			Handler:    _KaryawanService_ReadKaryawanByJabatan_Handler,
		},
		{
			MethodName: "ReadKaryawanByBagian",
			Handler:    _KaryawanService_ReadKaryawanByBagian_Handler,
		},
		{
			MethodName: "ReadKaryawanByKeterangan",
			Handler:    _KaryawanService_ReadKaryawanByKeterangan_Handler,
		},
		{
			MethodName: "ReadKaryawan",
			Handler:    _KaryawanService_ReadKaryawan_Handler,
		},
		{
			MethodName: "UpdateKaryawan",
			Handler:    _KaryawanService_UpdateKaryawan_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "karyawan.proto",
}

func init() { proto.RegisterFile("karyawan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0xfc, 0xcf, 0x24, 0x0d, 0xb0, 0x0a, 0xed, 0xe2, 0x96, 0x50, 0xf6, 0x94, 0x93,
	0x8b, 0x02, 0x48, 0xbd, 0x54, 0xa2, 0x69, 0x2a, 0x11, 0x8a, 0x88, 0x94, 0x82, 0x40, 0xdc, 0x26,
	0xc9, 0x12, 0x45, 0x38, 0x8e, 0x71, 0x5c, 0x50, 0x5e, 0x94, 0xf7, 0x40, 0xe2, 0xc2, 0x0d, 0xed,
	0xda, 0xb1, 0x77, 0x93, 0xac, 0xb9, 0x71, 0xea, 0x2d, 0x33, 0xdf, 0xcc, 0x64, 0xfd, 0xf3, 0x7e,
	0x1e, 0x68, 0x7c, 0xc5, 0x60, 0x85, 0x3f, 0xd0, 0x73, 0xfc, 0x60, 0x11, 0x2e, 0x48, 0x61, 0x1a,
	0xf8, 0x63, 0xfb, 0x68, 0xba, 0x58, 0x4c, 0x5d, 0x7e, 0x2a, 0x73, 0xa3, 0xdb, 0x2f, 0xa7, 0x7c,
	0xee, 0x87, 0xab, 0xa8, 0x84, 0xfd, 0xcc, 0x41, 0xe3, 0x62, 0x32, 0xb9, 0x8e, 0x1b, 0x87, 0xfc,
	0x1b, 0x21, 0x50, 0x78, 0x87, 0x73, 0xa4, 0xd6, 0x89, 0xd5, 0xae, 0x0e, 0xe5, 0x6f, 0x72, 0x00,
	0xa5, 0x0b, 0x17, 0xe7, 0x18, 0xd2, 0x9c, 0xcc, 0xc6, 0x11, 0xa1, 0x50, 0x7e, 0xcf, 0x5d, 0xee,
	0x2f, 0x3c, 0x9a, 0x97, 0xc2, 0x3a, 0x24, 0x4d, 0x28, 0x5e, 0xcd, 0x71, 0xe6, 0xd2, 0x82, 0xcc,
	0x47, 0x01, 0x39, 0x86, 0x6a, 0xbf, 0xf7, 0x06, 0x47, 0x18, 0xa2, 0x47, 0x8b, 0x52, 0x49, 0x13,
	0xc4, 0x86, 0x4a, 0xbf, 0xd7, 0xc5, 0xe9, 0x0c, 0x3d, 0x5a, 0x92, 0x62, 0x12, 0x8b, 0x13, 0xdc,
	0x84, 0x18, 0xde, 0x2e, 0x69, 0xf9, 0xc4, 0x6a, 0x17, 0x87, 0x71, 0x24, 0x26, 0x5e, 0x06, 0x1c,
	0x43, 0x3e, 0xe9, 0xae, 0x68, 0x25, 0x9a, 0x98, 0x24, 0x14, 0x75, 0xe0, 0xd1, 0xaa, 0xa6, 0x0e,
	0x3c, 0xa1, 0x7e, 0xf0, 0x27, 0x71, 0x2f, 0x44, 0x6a, 0x92, 0x50, 0xd4, 0x81, 0x47, 0x6b, 0x9a,
	0x3a, 0xf0, 0x48, 0x0b, 0xe0, 0x9a, 0x87, 0x3c, 0x40, 0x6f, 0x8a, 0x1e, 0xad, 0x4b, 0x59, 0xc9,
	0xb0, 0xdf, 0x39, 0xa8, 0x0f, 0x39, 0x26, 0x64, 0x45, 0x43, 0xbf, 0xb7, 0x8e, 0x62, 0xb8, 0x4a,
	0x26, 0xc1, 0x9e, 0xdb, 0x89, 0x3d, 0x6f, 0xc2, 0x5e, 0x30, 0x60, 0x2f, 0x1a, 0xb1, 0x97, 0xb2,
	0xb0, 0x97, 0x8d, 0xd8, 0x2b, 0x66, 0xec, 0xd5, 0x4c, 0xec, 0x90, 0x89, 0xbd, 0x96, 0x89, 0xbd,
	0x9e, 0x8d, 0x7d, 0x7f, 0x0b, 0xfb, 0x19, 0x50, 0x95, 0x7a, 0x77, 0x15, 0x3f, 0xa4, 0xb8, 0xd8,
	0x1a, 0x05, 0x6b, 0x83, 0x02, 0x1b, 0xc0, 0x23, 0x43, 0xe7, 0xd2, 0x27, 0x1d, 0x00, 0x74, 0xdd,
	0xb4, 0x37, 0xdf, 0xae, 0x75, 0x88, 0x23, 0xec, 0xe5, 0xa8, 0x4d, 0x43, 0xa5, 0x8a, 0xbd, 0x84,
	0x43, 0x7d, 0x60, 0x84, 0x54, 0x9c, 0x44, 0x25, 0x6e, 0xe9, 0xc4, 0xd9, 0xdb, 0xcd, 0x27, 0x58,
	0xb7, 0x2d, 0x7d, 0xf2, 0x0c, 0xaa, 0xe8, 0xba, 0x49, 0xa3, 0xe9, 0x14, 0x69, 0x11, 0x3b, 0x87,
	0x23, 0x7d, 0x5a, 0xca, 0x4a, 0x1c, 0x44, 0xc7, 0x69, 0x6d, 0xe1, 0xfc, 0x04, 0xc7, 0xe6, 0xf6,
	0xa5, 0x4f, 0xce, 0x60, 0x1f, 0x5d, 0x57, 0x1b, 0x61, 0x3a, 0x94, 0x5e, 0xc8, 0x5e, 0xc3, 0x7d,
	0x4d, 0x16, 0xd3, 0x5e, 0x40, 0x4d, 0x14, 0xa5, 0x1e, 0x31, 0xcd, 0x52, 0xcb, 0xd8, 0x9f, 0x1c,
	0x3c, 0x88, 0x2e, 0x88, 0xfa, 0x15, 0xbb, 0xb3, 0xdb, 0x7f, 0xb0, 0x5b, 0xe7, 0x57, 0x1e, 0xee,
	0xad, 0x91, 0xde, 0xf0, 0xe0, 0xfb, 0x6c, 0xcc, 0xc9, 0x39, 0xd4, 0x94, 0x8d, 0x42, 0x9a, 0xd1,
	0xfb, 0xd3, 0x97, 0x8c, 0x7d, 0xe0, 0x44, 0x5b, 0xc9, 0x59, 0x6f, 0x25, 0xe7, 0x4a, 0x6c, 0x25,
	0xb6, 0x47, 0x3e, 0xc3, 0xc3, 0x9d, 0x3e, 0x24, 0xad, 0xed, 0x8b, 0xa0, 0xda, 0xdb, 0x7e, 0x92,
	0xa9, 0x2f, 0x7d, 0xb6, 0x47, 0x3e, 0x42, 0x73, 0x97, 0xb7, 0xc8, 0xe3, 0x5d, 0xad, 0x89, 0x5d,
	0xed, 0x56, 0x96, 0x2c, 0x07, 0x8f, 0x37, 0x4d, 0x9b, 0x32, 0x22, 0x4f, 0x77, 0x75, 0x6b, 0x36,
	0xb4, 0xd9, 0xbf, 0x4a, 0xe4, 0x9f, 0xbc, 0xda, 0xd8, 0x28, 0x06, 0x86, 0x82, 0xed, 0x96, 0x63,
	0xa2, 0x09, 0x97, 0xd0, 0xd0, 0x9d, 0x42, 0x0e, 0xa3, 0xda, 0x2d, 0xff, 0x98, 0x5f, 0xd0, 0xa8,
	0x24, 0x33, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x60, 0x40, 0xb2, 0x7e, 0x6e, 0x08, 0x00,
	0x00,
}