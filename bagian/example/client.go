package main

import (
	"context"
	"time"

	cli "MiniProject/restoran/bagian/endpoint"
	svc "MiniProject/restoran/bagian/server"
	opt "MiniProject/restoran/util/grpc"

	util "MiniProject/restoran/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCBagianClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Bagian
	client.AddBagianService(context.Background(), svc.Bagian{NamaBagian: "", Deskripsi: "", Status: "", CreatedBy: "", CreatedOn: "", UpdatedBy: "", UpdateOn: ""})

	//Get Bagian Aktif
	//bagStat, _ := client.ReadBagianAktifService(context.Background())
	//fmt.Println("Bagian Aktif:",bagStat)

	//List Bagian
	//bags, _ := client.ReadBagianService(context.Background())
	//fmt.Println("all Bagians:",bags)

	//Update Bagian
	//client.UpdateBagianService(context.Background(), svc.Bagian{NamaBagian:"", Deskripsi:"", Status:"", CreatedBy:"", CreatedOn:"", UpdatedBy:"", UpdateOn:"",IDBagian:""})
}
