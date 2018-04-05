package endpoint

import (
	"context"
	"time"

	svc "MiniProject/restoran/karyawan/server"

	pb "MiniProject/restoran/karyawan/grpc"

	util "MiniProject/restoran/util/grpc"
	disc "MiniProject/restoran/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.KaryawanService"
)

func NewGRPCKaryawanClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KaryawanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKaryawanEp = retry
	}

	var readKaryawanByJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByJabatanEp = retry
	}

	var readKaryawanByBagianEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByBagian, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByBagianEp = retry
	}
	var readKaryawanByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByKeterangan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByKeteranganEp = retry
	}
	var readKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanEp = retry
	}

	var updateKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKaryawan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKaryawanEp = retry
	}

	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp, ReadKaryawanByJabatanEndpoint: readKaryawanByJabatanEp,
		ReadKaryawanByKeteranganEndpoint: readKaryawanByKeteranganEp, ReadKaryawanByBagianEndpoint: readKaryawanByBagianEp,
		ReadKaryawanEndpoint: readKaryawanEp, UpdateKaryawanEndpoint: updateKaryawanEp,
	}, nil
}

func encodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.AddKaryawanReq{
		Nama:       req.Nama,
		Alamat:     req.Alamat,
		Telepon:    req.Telepon,
		Email:      req.Email,
		IDJabatan:  req.IDJabatan,
		IDBagian:   req.IDBagian,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadKaryawanByJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByJabatanReq{IDJabatan: req.IDJabatan}, nil
}

func encodeReadKaryawanByBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByBagianReq{IDBagian: req.IDBagian}, nil
}

func encodeReadKaryawanByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.UpdateKaryawanReq{
		IDKaryawan: req.IDKaryawan,
		Nama:       req.Nama,
		Alamat:     req.Alamat,
		Telepon:    req.Telepon,
		Email:      req.Email,
		IDJabatan:  req.IDJabatan,
		IDBagian:   req.IDBagian,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn,
		Keterangan: req.Keterangan,
	}, nil
}

func decodeKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKaryawanByJabatanRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByJabatanResp)
	var rsp svc.Karyawans
	for _, v := range resp.AllJabatan {
		itm := svc.Karyawan{
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
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKaryawanByBagianRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByBagianResp)
	var rsp svc.Karyawans
	for _, v := range resp.AllBagian {
		itm := svc.Karyawan{
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
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKaryawanByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByKeteranganResp)
	var rsp svc.Karyawans
	for _, v := range resp.AllKeterangan {
		itm := svc.Karyawan{
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
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanResp)
	var rsp svc.Karyawans

	for _, v := range resp.AllKaryawan {
		itm := svc.Karyawan{
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
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKaryawan",
		encodeAddKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByJabatanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByJabatan",
		encodeReadKaryawanByJabatanRequest,
		decodeReadKaryawanByJabatanRespones,
		pb.ReadKaryawanByJabatanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByBagian(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByBagian",
		encodeReadKaryawanByBagianRequest,
		decodeReadKaryawanByBagianRespones,
		pb.ReadKaryawanByBagianResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByBagian")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByBagian",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByKeterangan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByKeterangan",
		encodeReadKaryawanByKeteranganRequest,
		decodeReadKaryawanByKeteranganRespones,
		pb.ReadKaryawanByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawan",
		encodeReadKaryawanRequest,
		decodeReadKaryawanResponse,
		pb.ReadKaryawanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKaryawan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKaryawan",
		encodeUpdateKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
