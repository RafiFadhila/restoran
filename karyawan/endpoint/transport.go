package endpoint

import (
	"context"

	scv "MiniProject/restoran/karyawan/server"
 
	pb "MiniProject/restoran/karyawan/grpc"
 
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKaryawanServer struct {
	addKaryawan              grpctransport.Handler
	readKaryawanByJabatan    grpctransport.Handler
	readKaryawanByBagian     grpctransport.Handler
	readKaryawanByKeterangan grpctransport.Handler
	readKaryawan             grpctransport.Handler
	updateKaryawan           grpctransport.Handler
}

func NewGRPCKaryawanServer(endpoints KaryawanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KaryawanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKaryawanServer{
		addKaryawan: grpctransport.NewServer(endpoints.AddKaryawanEndpoint,
			decodeAddKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKaryawan", logger)))...),
		readKaryawanByJabatan: grpctransport.NewServer(endpoints.ReadKaryawanByJabatanEndpoint,
			decodeReadKaryawanByJabatanRequest,
			encodeReadKaryawanByJabatanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByJabatan", logger)))...),
		readKaryawanByBagian: grpctransport.NewServer(endpoints.ReadKaryawanByBagianEndpoint,
			decodeReadKaryawanByBagianRequest,
			encodeReadKaryawanByBagianResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByBagian", logger)))...),
		readKaryawanByKeterangan: grpctransport.NewServer(endpoints.ReadKaryawanByKeteranganEndpoint,
			decodeReadKaryawanByKeteranganRequest,
			encodeReadKaryawanByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByKeterangan", logger)))...),
		readKaryawan: grpctransport.NewServer(endpoints.ReadKaryawanEndpoint,
			decodeReadKaryawanRequest,
			encodeReadKaryawanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawan", logger)))...),
		updateKaryawan: grpctransport.NewServer(endpoints.UpdateKaryawanEndpoint,
			decodeUpdateKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKaryawan", logger)))...),
	}
}

func decodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKaryawanReq)
	return scv.Karyawan{Nama: req.GetNama(), Alamat: req.GetAlamat(), Telepon: req.GetTelepon(), Email: req.GetEmail(), IDJabatan: req.GetIDJabatan(), IDBagian: req.GetIDBagian(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdatedBy: req.GetUpdatedBy(), UpdatedOn: req.GetUpdatedOn(), Keterangan: req.GetKeterangan()}, nil
}

func decodeReadKaryawanByJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByJabatanReq)
	return scv.Karyawan{IDJabatan: req.IDJabatan}, nil
}
func decodeReadKaryawanByBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByBagianReq)
	return scv.Karyawan{IDBagian: req.IDBagian}, nil

}
func decodeReadKaryawanByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByKeteranganReq)
	return scv.Karyawan{Keterangan: req.Keterangan}, nil

}

func decodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKaryawanReq)
	return scv.Karyawan{IDKaryawan: req.IDKaryawan, Nama: req.Nama, Alamat: req.Alamat, Telepon: req.Telepon, Email: req.Email, IDJabatan: req.IDJabatan, IDBagian: req.IDBagian, Status: req.Status, CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdatedBy: req.UpdatedBy, UpdatedOn: req.UpdatedOn, Keterangan: req.Keterangan}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadKaryawanByJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanByJabatanResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawan{
			IDKaryawan: v.IDKaryawan,
			Nama:       v.Nama,
			Alamat:     v.Alamat,
			Telepon:    v.Telepon,
			Email:      v.Email,
			IDJabatan:  v.IDJabatan,
			IDBagian:   v.IDBagian,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdatedOn:  v.UpdatedOn,
			Keterangan: v.Keterangan,
		}
		rsp.AllJabatan = append(rsp.AllJabatan, itm)
	}
	return rsp, nil
}

func encodeReadKaryawanByBagianResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanByBagianResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawan{
			IDKaryawan: v.IDKaryawan,
			Nama:       v.Nama,
			Alamat:     v.Alamat,
			Telepon:    v.Telepon,
			Email:      v.Email,
			IDJabatan:  v.IDJabatan,
			IDBagian:   v.IDBagian,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdatedOn:  v.UpdatedOn,
			Keterangan: v.Keterangan,
		}
		rsp.AllBagian = append(rsp.AllBagian, itm)
	}
	return rsp, nil
}

func encodeReadKaryawanByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawan{
			IDKaryawan: v.IDKaryawan,
			Nama:       v.Nama,
			Alamat:     v.Alamat,
			Telepon:    v.Telepon,
			Email:      v.Email,
			IDJabatan:  v.IDJabatan,
			IDBagian:   v.IDBagian,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdatedOn:  v.UpdatedOn,
			Keterangan: v.Keterangan,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func encodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawan{
			IDKaryawan: v.IDKaryawan,
			Nama:       v.Nama,
			Alamat:     v.Alamat,
			Telepon:    v.Telepon,
			Email:      v.Email,
			IDJabatan:  v.IDJabatan,
			IDBagian:   v.IDBagian,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdatedOn:  v.UpdatedOn,
			Keterangan: v.Keterangan,
		}
		rsp.AllKaryawan = append(rsp.AllKaryawan, itm)
	}
	return rsp, nil
}

func (s *grpcKaryawanServer) AddKaryawan(ctx oldcontext.Context, karyawan *pb.AddKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKaryawan.ServeGRPC(ctx, karyawan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByJabatan(ctx oldcontext.Context, jabatan *pb.ReadKaryawanByJabatanReq) (*pb.ReadKaryawanByJabatanResp, error) {
	_, resp, err := s.readKaryawanByJabatan.ServeGRPC(ctx, jabatan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByJabatanResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByBagian(ctx oldcontext.Context, bagian *pb.ReadKaryawanByBagianReq) (*pb.ReadKaryawanByBagianResp, error) {
	_, resp, err := s.readKaryawanByBagian.ServeGRPC(ctx, bagian)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByBagianResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawanByKeterangan(ctx oldcontext.Context, ket *pb.ReadKaryawanByKeteranganReq) (*pb.ReadKaryawanByKeteranganResp, error) {
	_, resp, err := s.readKaryawanByKeterangan.ServeGRPC(ctx, ket)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByKeteranganResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKaryawanResp, error) {
	_, resp, err := s.readKaryawan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanResp), nil
}

func (s *grpcKaryawanServer) UpdateKaryawan(ctx oldcontext.Context, cus *pb.UpdateKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKaryawan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
