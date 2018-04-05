package endpoint

import (
	"context"

	svc "MiniProject/restoran/jabatan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type JabatanEndpoint struct {
	AddJabatanEndpoint       kit.Endpoint
	ReadJabatanAktifEndpoint kit.Endpoint
	ReadJabatanEndpoint      kit.Endpoint
	UpdateJabatanEndpoint    kit.Endpoint
}

func NewJabatanEndpoint(service svc.JabatanService) JabatanEndpoint {
	addJabatanEp := makeAddJabatanEndpoint(service)
	readJabatanAktifEp := makeReadJabatanAktifEndpoint(service)
	readJabatanEp := makeReadJabatanEndpoint(service)
	updateJabatanEp := makeUpdateJabatanEndpoint(service)
	return JabatanEndpoint{AddJabatanEndpoint: addJabatanEp,
		ReadJabatanAktifEndpoint: readJabatanAktifEp,
		ReadJabatanEndpoint:      readJabatanEp,
		UpdateJabatanEndpoint:    updateJabatanEp,
	}
}

func makeAddJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Jabatan)
		err := service.AddJabatanService(ctx, req)
		return nil, err
	}
}

func makeReadJabatanAktifEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadJabatanAktifService(ctx)
		return result, err
	}
}

func makeReadJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadJabatanService(ctx)
		return result, err
	}
}

func makeUpdateJabatanEndpoint(service svc.JabatanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Jabatan)
		err := service.UpdateJabatanService(ctx, req)
		return nil, err
	}
}
