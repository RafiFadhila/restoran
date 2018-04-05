package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/restoran/jabatan/server"
)

func (ke JabatanEndpoint) AddJabatanService(ctx context.Context, jabatan sv.Jabatan) error {
	_, err := ke.AddJabatanEndpoint(ctx, jabatan)
	return err
}

func (ce JabatanEndpoint) ReadJabatanAktifService(ctx context.Context) (sv.Jabatans, error) {
	resp, err := ce.ReadJabatanAktifEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Jabatans), err
}

func (ce JabatanEndpoint) ReadJabatanService(ctx context.Context) (sv.Jabatans, error) {
	resp, err := ce.ReadJabatanEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Jabatans), err
}

func (ce JabatanEndpoint) UpdateJabatanService(ctx context.Context, cus sv.Jabatan) error {
	_, err := ce.UpdateJabatanEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
