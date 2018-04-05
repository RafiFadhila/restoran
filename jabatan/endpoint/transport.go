package endpoint

import (
	"context"

	scv "MiniProject/restoran/jabatan/server"

	pb "MiniProject/restoran/jabatan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcJabatanServer struct {
	addJabatan       grpctransport.Handler
	readJabatanAktif grpctransport.Handler
	readJabatan      grpctransport.Handler
	updateJabatan    grpctransport.Handler
}

func NewGRPCJabatanServer(endpoints JabatanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.JabatanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcJabatanServer{
		addJabatan: grpctransport.NewServer(endpoints.AddJabatanEndpoint,
			decodeAddJabatanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddJabatan", logger)))...),
		readJabatanAktif: grpctransport.NewServer(endpoints.ReadJabatanAktifEndpoint,
			decodeReadJabatanAktifRequest,
			encodeReadJabatanAktifResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJabatanAktif", logger)))...),
		readJabatan: grpctransport.NewServer(endpoints.ReadJabatanEndpoint,
			decodeReadJabatanRequest,
			encodeReadJabatanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadJabatan", logger)))...),
		updateJabatan: grpctransport.NewServer(endpoints.UpdateJabatanEndpoint,
			decodeUpdateJabatanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateJabatan", logger)))...),
	}
}

func decodeAddJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddJabatanReq)
	return scv.Jabatan{Jabatan: req.GetJabatan(), Gaji: req.GetGaji(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdatedBy: req.GetUpdatedBy(), UpdateOn: req.GetUpdateOn()}, nil
}

func decodeReadJabatanAktifRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJabatanReq)
	return scv.Jabatan{IDJabatan: req.IDJabatan, Jabatan: req.Jabatan, Gaji: req.Gaji, Status: req.Status, CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdatedBy: req.UpdatedBy, UpdateOn: req.UpdateOn}, nil

}
func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadJabatanAktifResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Jabatans)

	rsp := &pb.ReadJabatanAktifResp{}

	for _, v := range resp {
		itm := &pb.ReadJabatan{
			IDJabatan: v.IDJabatan,
			Jabatan:   v.Jabatan,
			Gaji:      v.Gaji,
			Status:    v.Status,
			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdatedBy: v.UpdatedBy,
			UpdateOn:  v.UpdateOn,
		}
		rsp.AllStatus = append(rsp.AllStatus, itm)
	}
	return rsp, nil
}

func encodeReadJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Jabatans)

	rsp := &pb.ReadJabatanResp{}

	for _, v := range resp {
		itm := &pb.ReadJabatan{
			IDJabatan: v.IDJabatan,
			Jabatan:   v.Jabatan,
			Gaji:      v.Gaji,
			Status:    v.Status,
			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdatedBy: v.UpdatedBy,
			UpdateOn:  v.UpdateOn,
		}
		rsp.AllJabatan = append(rsp.AllJabatan, itm)
	}
	return rsp, nil
}

func (s *grpcJabatanServer) AddJabatan(ctx oldcontext.Context, jabatan *pb.AddJabatanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addJabatan.ServeGRPC(ctx, jabatan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcJabatanServer) ReadJabatanAktif(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadJabatanAktifResp, error) {
	_, resp, err := s.readJabatanAktif.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJabatanAktifResp), nil
}

func (s *grpcJabatanServer) ReadJabatan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadJabatanResp, error) {
	_, resp, err := s.readJabatan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadJabatanResp), nil
}

func (s *grpcJabatanServer) UpdateJabatan(ctx oldcontext.Context, cus *pb.UpdateJabatanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateJabatan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
