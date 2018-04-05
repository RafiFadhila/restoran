package endpoint

import (
	"context"
	"time"

	svc "MiniProject/restoran/bagian/server"

	pb "MiniProject/restoran/bagian/grpc"

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
	grpcName = "grpc.BagianService"
)

func NewGRPCBagianClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.BagianService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addBagianEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddBagianEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addBagianEp = retry
	}

	var readBagianAktifEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBagianAktifEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBagianAktifEp = retry
	}

	var readBagianEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadBagianEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readBagianEp = retry
	}

	var updateBagianEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateBagian, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateBagianEp = retry
	}
	return BagianEndpoint{AddBagianEndpoint: addBagianEp, ReadBagianAktifEndpoint: readBagianAktifEp,
		ReadBagianEndpoint: readBagianEp, UpdateBagianEndpoint: updateBagianEp}, nil
}

func encodeAddBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Bagian)
	return &pb.AddBagianReq{
		NamaBagian: req.NamaBagian,
		Deskripsi:  req.Deskripsi,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdateOn:   req.UpdateOn,
	}, nil
}

func encodeReadBagianAktifRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateBagianRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Bagian)
	return &pb.UpdateBagianReq{
		IDBagian:   req.IDBagian,
		NamaBagian: req.NamaBagian,
		Deskripsi:  req.Deskripsi,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdateOn:   req.UpdateOn,
	}, nil
}

func decodeBagianResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadBagianAktifResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBagianAktifResp)
	var rsp svc.Bagians

	for _, v := range resp.AllAktif {
		itm := svc.Bagian{
			IDBagian:   v.IDBagian,
			NamaBagian: v.NamaBagian,
			Deskripsi:  v.Deskripsi,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdateOn:   v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadBagianResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadBagianResp)
	var rsp svc.Bagians

	for _, v := range resp.AllBagian {
		itm := svc.Bagian{
			IDBagian:   v.IDBagian,
			NamaBagian: v.NamaBagian,
			Deskripsi:  v.Deskripsi,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdatedBy:  v.UpdatedBy,
			UpdateOn:   v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddBagianEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddBagian",
		encodeAddBagianRequest,
		decodeBagianResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddBagian")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddBagian",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBagianAktifEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBagianAktif",
		encodeReadBagianAktifRequest,
		decodeReadBagianAktifResponse,
		pb.ReadBagianAktifResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBagianAktif")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBagianAktif",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadBagianEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadBagian",
		encodeReadBagianRequest,
		decodeReadBagianResponse,
		pb.ReadBagianResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadBagian")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadBagian",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateBagian(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateBagian",
		encodeUpdateBagianRequest,
		decodeBagianResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKaryawam")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateBagian",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
