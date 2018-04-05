package server

import (
	"context"
)

type jabatan struct {
	writer ReadWriter
}

func NewJabatan(writer ReadWriter) JabatanService {
	return &jabatan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (k *jabatan) AddJabatanService(ctx context.Context, jabatan Jabatan) error {
	//fmt.Println("customer")
	err := k.writer.AddJabatan(jabatan)
	if err != nil {
		return err
	}

	return nil
}

func (k *jabatan) ReadJabatanAktifService(ctx context.Context) (Jabatans, error) {
	kar, err := k.writer.ReadJabatan()
	//fmt.Println("customer", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *jabatan) ReadJabatanService(ctx context.Context) (Jabatans, error) {
	kar, err := k.writer.ReadJabatan()
	//fmt.Println("customer", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *jabatan) UpdateJabatanService(ctx context.Context, kar Jabatan) error {
	err := k.writer.UpdateJabatan(kar)
	if err != nil {
		return err
	}
	return nil
}
