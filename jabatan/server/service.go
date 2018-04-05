package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "jabatan.restoran.id"
	OnAdd     Status = 1
)

type Jabatan struct {
	IDJabatan string
	Jabatan   string
	Gaji      int64
	Status    int32
	CreatedBy string
	CreatedOn string
	UpdatedBy string
	UpdateOn  string
}
type Jabatans []Jabatan

type ReadWriter interface {
	AddJabatan(Jabatan) error
	ReadJabatanAktif() (Jabatans, error)
	ReadJabatan() (Jabatans, error)
	UpdateJabatan(Jabatan) error
}

type JabatanService interface {
	AddJabatanService(context.Context, Jabatan) error
	ReadJabatanAktifService(context.Context) (Jabatans, error)
	ReadJabatanService(context.Context) (Jabatans, error)
	UpdateJabatanService(context.Context, Jabatan) error
}
