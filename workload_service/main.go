package workloadservice

import (
	"log"
	"net"

	"RailwayStationsWorkload_micro/workload_service/protobuff"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

const WLport = ":9092"

func NewWorkload() {
	logg := hclog.Default()
	lis, err := net.Listen("tcp", WLport)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	gs := grpc.NewServer()
	ms := NewMyServer(logg)
	protobuff.RegisterWorkloadServiceServer(gs, ms)
	//reflection.Register(reflection.GRPCServer(gs))
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9092")
	go func() {
		log.Fatalln(gs.Serve(lis))
	}()
}
