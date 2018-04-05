package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "karyawan.restoran.id"
	OnAdd     Status = 1
)

type Karyawan struct {
	IDKaryawan string
	Nama       string
	Alamat     string
	Telepon    string
	Email      string
	IDJabatan  string
	IDBagian   string
	Status     int32
	CreatedBy  string
	CreatedOn  string
	UpdatedBy  string
	UpdatedOn  string
	Keterangan string
}
type Karyawans []Karyawan

type ReadWriter interface {
	AddKaryawan(Karyawan) error
	ReadKaryawanByBagian(string) (Karyawans, error)
	ReadKaryawanByJabatan(string) (Karyawans, error)
	ReadKaryawanByKeterangan(string) (Karyawans, error)
	ReadKaryawan() (Karyawans, error)
	UpdateKaryawan(Karyawan) error
}

type KaryawanService interface {
	AddKaryawanService(context.Context, Karyawan) error
	ReadKaryawanByBagianService(context.Context, string) (Karyawans, error)
	ReadKaryawanByJabatanService(context.Context, string) (Karyawans, error)
	ReadKaryawanByKeteranganService(context.Context, string) (Karyawans, error)
	ReadKaryawanService(context.Context) (Karyawans, error)
	UpdateKaryawanService(context.Context, Karyawan) error
}
