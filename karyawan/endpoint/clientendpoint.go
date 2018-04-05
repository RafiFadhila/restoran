package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/restoran/karyawan/server"
)

func (ke KaryawanEndpoint) AddKaryawanService(ctx context.Context, karyawan sv.Karyawan) error {
	_, err := ke.AddKaryawanEndpoint(ctx, karyawan)
	return err
}

func (ce KaryawanEndpoint) ReadKaryawanByJabatanService(ctx context.Context, jabatan string) (sv.Karyawans, error) {
	req := sv.Karyawan{IDJabatan: jabatan}
	fmt.Println(req)
	resp, err := ce.ReadKaryawanByJabatanEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Karyawans)
	return cus, err
}

func (ce KaryawanEndpoint) ReadKaryawanByBagianService(ctx context.Context, bagian string) (sv.Karyawans, error) {
	req := sv.Karyawan{IDBagian: bagian}
	resp, err := ce.ReadKaryawanByBagianEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Karyawans)
	return cus, err
}

func (ce KaryawanEndpoint) ReadKaryawanByKeteranganService(ctx context.Context, keterangan string) (sv.Karyawans, error) {
	req := sv.Karyawan{Keterangan: keterangan}
	resp, err := ce.ReadKaryawanByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Karyawans)
	return cus, err
}

func (ce KaryawanEndpoint) ReadKaryawanService(ctx context.Context) (sv.Karyawans, error) {
	resp, err := ce.ReadKaryawanEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Karyawans), err
}

func (ce KaryawanEndpoint) UpdateKaryawanService(ctx context.Context, cus sv.Karyawan) error {
	_, err := ce.UpdateKaryawanEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
