package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "bagian.restoran.id"
	OnAdd     Status = 1
)

type Bagian struct {
	IDBagian   string
	NamaBagian string
	Deskripsi  string
	Status     int32
	CreatedBy  string
	CreatedOn  string
	UpdatedBy  string
	UpdateOn   string
}
type Bagians []Bagian

type ReadWriter interface {
	AddBagian(Bagian) error
	ReadBagianAktif() (Bagians, error)
	ReadBagian() (Bagians, error)
	UpdateBagian(Bagian) error
}

type BagianService interface {
	AddBagianService(context.Context, Bagian) error
	ReadBagianAktifService(context.Context) (Bagians, error)
	ReadBagianService(context.Context) (Bagians, error)
	UpdateBagianService(context.Context, Bagian) error
}
