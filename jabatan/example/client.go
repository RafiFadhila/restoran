package main

import (
	"context"
	"time"

	cli "MiniProject/restoran/jabatan/endpoint"
	svc "MiniProject/restoran/jabatan/server"
	opt "MiniProject/restoran/util/grpc"

	util "MiniProject/restoran/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCJabatanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Jabatan
	client.AddJabatanService(context.Background(), svc.Jabatan{Jabatan:"", Gaji:, Status:, CreatedBy:"", CreatedOn:"", UpdatedBy:"", UpdateOn:""})

	//Get Jabatan Aktif
	//jabStat, _ := client.ReadJabatanAktifService(context.Background())
	//fmt.Println("Jabatan Aktif:",jabStat)

	//List Aktif
	//jabs, _ := client.ReadJabatanService(context.Background())
	//fmt.Println("all Jabatans:",jabs)

	//Update Jabatan
	//client.UpdateJabatanService(context.Background(), svc.Jabatan{Jabatan:"", Gaji:, Status:, CreatedBy:"", CreatedOn:"", UpdatedBy:"", UpdateOn:"", IDJabatan:})
}
