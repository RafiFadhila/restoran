package server

import (
	"context"
)

type karyawan struct {
	writer ReadWriter
}

func NewKaryawan(writer ReadWriter) KaryawanService {
	return &karyawan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (k *karyawan) AddKaryawanService(ctx context.Context, karyawan Karyawan) error {
	//fmt.Println("customer")
	err := k.writer.AddKaryawan(karyawan)
	if err != nil {
		return err
	}

	return nil
}

func (k *karyawan) ReadKaryawanByJabatanService(ctx context.Context, jab string) (Karyawans, error) {
	kar, err := k.writer.ReadKaryawanByJabatan(jab)
	//fmt.Println(cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *karyawan) ReadKaryawanByBagianService(ctx context.Context, bag string) (Karyawans, error) {
	cus, err := k.writer.ReadKaryawanByBagian(bag)
	//fmt.Println("customer:", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (k *karyawan) ReadKaryawanByKeteranganService(ctx context.Context, ket string) (Karyawans, error) {
	cus, err := k.writer.ReadKaryawanByKeterangan(ket)
	//fmt.Println("customer:", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (k *karyawan) ReadKaryawanService(ctx context.Context) (Karyawans, error) {
	kar, err := k.writer.ReadKaryawan()
	//fmt.Println("customer", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *karyawan) UpdateKaryawanService(ctx context.Context, kar Karyawan) error {
	err := k.writer.UpdateKaryawan(kar)
	if err != nil {
		return err
	}
	return nil
}
