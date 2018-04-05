package endpoint

import (
	"context"

	svc "MiniProject/restoran/bagian/server"

	kit "github.com/go-kit/kit/endpoint"
)

type BagianEndpoint struct {
	AddBagianEndpoint       kit.Endpoint
	ReadBagianAktifEndpoint kit.Endpoint
	ReadBagianEndpoint      kit.Endpoint
	UpdateBagianEndpoint    kit.Endpoint
}

func NewBagianEndpoint(service svc.BagianService) BagianEndpoint {
	addBagianEp := makeAddBagianEndpoint(service)
	readBagianAktifEp := makeReadBagianAktifEndpoint(service)
	readBagianEp := makeReadBagianEndpoint(service)
	updateBagianEp := makeUpdateBagianEndpoint(service)
	return BagianEndpoint{AddBagianEndpoint: addBagianEp,
		ReadBagianAktifEndpoint: readBagianAktifEp,
		ReadBagianEndpoint:      readBagianEp,
		UpdateBagianEndpoint:    updateBagianEp,
	}
}

func makeAddBagianEndpoint(service svc.BagianService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Bagian)
		err := service.AddBagianService(ctx, req)
		return nil, err
	}
}

func makeReadBagianAktifEndpoint(service svc.BagianService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadBagianAktifService(ctx)
		return result, err
	}
}

func makeReadBagianEndpoint(service svc.BagianService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadBagianService(ctx)
		return result, err
	}
}

func makeUpdateBagianEndpoint(service svc.BagianService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Bagian)
		err := service.UpdateBagianService(ctx, req)
		return nil, err
	}
}
