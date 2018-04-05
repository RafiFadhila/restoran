package endpoint

import (
	"context"
	"time"

	svc "MiniProject/restoran/jabatan/server"

	pb "MiniProject/restoran/jabatan/grpc"

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
	grpcName = "grpc.jabatanService"
)

func NewGRPCCustomerClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.JabatanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addJabatanEp = retry
	}
	var readJabatanAktifEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJabatanAktifEp = retry
	}
	var readJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadJabatanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readJabatanEp = retry
	}

	var updateJabatanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateJabatan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateJabatanEp = retry
	}
	return JabatanEndpoint{AddJabatanEndpoint: addJabatanEp, ReadJabatanAktifEndpoint: readJabatanAktifEp,
		ReadJabatanEndpoint: readJabatanEp, UpdateJabatanEndpoint: updateJabatanEp}, nil
}

func encodeAddJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Jabatan)
	return &pb.AddJabatanReq{
		Jabatan:   req.Jabatan,
		Gaji:      req.Gaji,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn,
		UpdatedBy: req.UpdatedBy,
		UpdateOn:  req.UpdateOn,
	}, nil
}

func encodeReadJabatanAktifRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateJabatanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Jabatan)
	return &pb.UpdateJabatanReq{
		IDJabatan: req.IDJabatan,
		Jabatan:   req.Jabatan,
		Gaji:      req.Gaji,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn,
		UpdatedBy: req.UpdatedBy,
		UpdateOn:  req.UpdateOn,
	}, nil
}

func decodeJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadJabatanAktifResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJabatanAktifResp)
	var rsp svc.Jabatans

	for _, v := range resp.AllStatus {
		itm := svc.Jabatan{
			IDJabatan: v.IDJabatan,
			Jabatan:   v.Jabatan,
			Gaji:      v.Gaji,
			Status:    v.Status,
			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdatedBy: v.UpdatedBy,
			UpdateOn:  v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadJabatanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadJabatanResp)
	var rsp svc.Jabatans

	for _, v := range resp.AllJabatan {
		itm := svc.Jabatan{
			IDJabatan: v.IDJabatan,
			Jabatan:   v.Jabatan,
			Gaji:      v.Gaji,
			Status:    v.Status,
			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdatedBy: v.UpdatedBy,
			UpdateOn:  v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}
func makeClientAddJabatanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddJabatan",
		encodeAddJabatanRequest,
		decodeJabatanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJabatanAktifEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJabatanAktif",
		encodeReadJabatanRequest,
		decodeReadJabatanResponse,
		pb.ReadJabatanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJabatanAktif")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJabatanAktif",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadJabatanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadJabatan",
		encodeReadJabatanRequest,
		decodeReadJabatanResponse,
		pb.ReadJabatanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadJabatan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateJabatan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateJabatan",
		encodeUpdateJabatanRequest,
		decodeJabatanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKaryawam")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateJabatan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
