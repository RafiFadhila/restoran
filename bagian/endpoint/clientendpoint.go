package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/restoran/bagian/server"
)

func (ke BagianEndpoint) AddBagianService(ctx context.Context, bagian sv.Bagian) error {
	_, err := ke.AddBagianEndpoint(ctx, bagian)
	return err
}

func (ce BagianEndpoint) ReadBagianAktifService(ctx context.Context) (sv.Bagians, error) {
	resp, err := ce.ReadBagianAktifEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Bagians), err
}

func (ce BagianEndpoint) ReadBagianService(ctx context.Context) (sv.Bagians, error) {
	resp, err := ce.ReadBagianEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Bagians), err
}

func (ce BagianEndpoint) UpdateBagianService(ctx context.Context, cus sv.Bagian) error {
	_, err := ce.UpdateBagianEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
