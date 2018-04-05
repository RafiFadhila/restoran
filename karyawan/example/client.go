package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/restoran/karyawan/endpoint"
	opt "MiniProject/restoran/util/grpc"
	util "MiniProject/restoran/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKaryawanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Karyawan
	//	client.AddKaryawanService(context.Background(), svc.Karyawan{Nama: "Bunga", Alamat: "Bogor", Telepon: "087789917477", Email: "bunga@gmail.com", IDJabatan: "1", IDBagian: "1",Keterangan:""})

	//Get Karyawan By Jabatan
	//karJabatan, _ := client.ReadKaryawanByJabatanService(context.Background(), "1")
	//fmt.Println("Karyawan based on Bagian:", karJabatan)

	//Get Karyawan By Bagian
	//karBagian, _ := client.ReadKaryawanByBagianService(context.Background(), "1")
	//fmt.Println("Karyawan based on Bagian:", karBagian)

	//Get Karyawan By Keterangan
	karKet, _ := client.ReadKaryawanByKeteranganService(context.Background(), "%cu%")
	fmt.Println("Karyawan based on Keterangan:", karKet)

	//List Karyawan
	//kars, _ := client.ReadKaryawanService(context.Background())
	//fmt.Println("all Karyawans: ", kars)

	//Update Karyawan
	//client.UpdateKaryawanService(context.Background(), svc.Karyawan{Nama: "Bunga", Alamat: "Bekasi", Telepon: "087789917477", Email: "bunga@gmail.com", IDJabatan: "1", IDBagian: "1", Status: 1, CreatedBy: "1", UpdatedBy: "1", Keterangan: "Cuti", IDKaryawan: "3"})

}
