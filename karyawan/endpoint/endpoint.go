package endpoint

import (
	"context"

	svc "MiniProject/restoran/karyawan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KaryawanEndpoint struct {
	AddKaryawanEndpoint              kit.Endpoint
	ReadKaryawanByJabatanEndpoint    kit.Endpoint
	ReadKaryawanByBagianEndpoint     kit.Endpoint
	ReadKaryawanByKeteranganEndpoint kit.Endpoint
	ReadKaryawanEndpoint             kit.Endpoint
	UpdateKaryawanEndpoint           kit.Endpoint
}

func NewKaryawanEndpoint(service svc.KaryawanService) KaryawanEndpoint {
	addKaryawanEp := makeAddKaryawanEndpoint(service)
	readKaryawanByJabatanEp := makeReadKaryawanByJabatanEndpoint(service)
	readKaryawanByBagianEp := makeReadKaryawanByBagianEndpoint(service)
	readKaryawanByKeteranganEp := makeReadKaryawanByKeteranganEndpoint(service)
	readKaryawanEp := makeReadKaryawanEndpoint(service)
	updateKaryawanEp := makeUpdateKaryawanEndpoint(service)
	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp,
		ReadKaryawanByJabatanEndpoint:    readKaryawanByJabatanEp,
		ReadKaryawanByBagianEndpoint:     readKaryawanByBagianEp,
		ReadKaryawanByKeteranganEndpoint: readKaryawanByKeteranganEp,
		ReadKaryawanEndpoint:             readKaryawanEp,
		UpdateKaryawanEndpoint:           updateKaryawanEp,
	}
}

func makeAddKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.AddKaryawanService(ctx, req)
		return nil, err
	}
}

func makeReadKaryawanByJabatanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByJabatanService(ctx, req.IDJabatan)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadKaryawanByBagianEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByBagianService(ctx, req.IDBagian)
		return result, err
	}
}

func makeReadKaryawanByKeteranganEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}

func makeReadKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKaryawanService(ctx)
		return result, err
	}
}

func makeUpdateKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.UpdateKaryawanService(ctx, req)
		return nil, err
	}
}
