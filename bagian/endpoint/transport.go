package endpoint

import (
	"context"

	scv "MiniProject/restoran/bagian/server"

	pb "MiniProject/restoran/bagian/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcBagianServer struct {
	addBagian       grpctransport.Handler
	readBagianAktif grpctransport.Handler
	readBagian      grpctransport.Handler
	updateBagian    grpctransport.Handler
}

func NewGRPCBagianServer(endpoints BagianEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.BagianServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcBagianServer{
		addBagian: grpctransport.NewServer(endpoints.AddBagianEndpoint,
			decodeAddBagianRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddBagian", logger)))...),
		readBagianAktif: grpctransport.NewServer(endpoints.ReadBagianAktifEndpoint,
			decodeReadBagianAktifRequest,
			encodeReadBagianAktifResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBagianAktif", logger)))...),
		readBagian: grpctransport.NewServer(endpoints.ReadBagianEndpoint,
			decodeReadBagianRequest,
			encodeReadBagianResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadBagian", logger)))...),
		updateBagian: grpctransport.NewServer(endpoints.UpdateBagianEndpoint,
			decodeUpdateBagianRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateBagian", logger)))...),
	}
}

func decodeAddBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddBagianReq)
	return scv.Bagian{NamaBagian: req.GetNamaBagian(), Deskripsi: req.GetDeskripsi(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdatedBy: req.GetUpdatedBy(), UpdateOn: req.GetUpdateOn()}, nil
}

func decodeReadBagianAktifRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateBagianReq)
	return scv.Bagian{IDBagian: req.IDBagian, NamaBagian: req.NamaBagian, Deskripsi: req.Deskripsi, Status: req.Status, CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdatedBy: req.UpdatedBy, UpdateOn: req.UpdateOn}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadBagianAktifResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Bagians)

	rsp := &pb.ReadBagianAktifResp{}

	for _, v := range resp {
		itm := &pb.ReadBagian{
			IDBagian:   v.IDBagian,
			NamaBagian: v.NamaBagian,
			Deskripsi:  v.Deskripsi,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdateOn:   v.UpdateOn}
		rsp.AllAktif = append(rsp.AllAktif, itm)
	}
	return rsp, nil
}

func encodeReadBagianResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Bagians)

	rsp := &pb.ReadBagianResp{}

	for _, v := range resp {
		itm := &pb.ReadBagian{
			IDBagian:   v.IDBagian,
			NamaBagian: v.NamaBagian,
			Deskripsi:  v.Deskripsi,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdateOn:   v.UpdateOn}
		rsp.AllBagian = append(rsp.AllBagian, itm)
	}
	return rsp, nil
}

func (s *grpcBagianServer) AddBagian(ctx oldcontext.Context, bagian *pb.AddBagianReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addBagian.ServeGRPC(ctx, bagian)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcBagianServer) ReadBagianAktif(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadBagianAktifResp, error) {
	_, resp, err := s.readBagianAktif.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBagianAktifResp), nil
}

func (s *grpcBagianServer) ReadBagian(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadBagianResp, error) {
	_, resp, err := s.readBagian.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadBagianResp), nil
}

func (s *grpcBagianServer) UpdateBagian(ctx oldcontext.Context, cus *pb.UpdateBagianReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateBagian.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
