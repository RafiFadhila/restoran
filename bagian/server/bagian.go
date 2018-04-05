package server

import (
	"context"
)

type bagian struct {
	writer ReadWriter
}

func NewBagian(writer ReadWriter) BagianService {
	return &bagian{writer: writer}
}

//Methode pada interface BagianService di service.go
func (k *bagian) AddBagianService(ctx context.Context, bagian Bagian) error {
	//fmt.Println("Bagian")
	err := k.writer.AddBagian(bagian)
	if err != nil {
		return err
	}

	return nil
}

func (k *bagian) ReadBagianAktifService(ctx context.Context) (Bagians, error) {
	kar, err := k.writer.ReadBagianAktif()
	//fmt.Println("Bagian", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *bagian) ReadBagianService(ctx context.Context) (Bagians, error) {
	kar, err := k.writer.ReadBagian()
	//fmt.Println("Bagian", cus)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (k *bagian) UpdateBagianService(ctx context.Context, kar Bagian) error {
	err := k.writer.UpdateBagian(kar)
	if err != nil {
		return err
	}
	return nil
}
