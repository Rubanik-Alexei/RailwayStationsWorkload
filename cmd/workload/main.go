package gateway

import (
	"log"
	"net"
	"os"

	_ "RailwayStationsWorkload_micro/config"
	workloadservice "RailwayStationsWorkload_micro/internal/workload_service"
	"RailwayStationsWorkload_micro/internal/workload_service/protobuff"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func NewWorkload() {
	logg := hclog.Default()
	//lis, err := net.Listen("tcp", config.WLport)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9092"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	gs := grpc.NewServer()
	ms := workloadservice.NewMyServer(logg)
	protobuff.RegisterWorkloadServiceServer(gs, ms)
	//reflection.Register(reflection.GRPCServer(gs))
	// Serve gRPC server
	log.Println("Serving WL_gRPC on 0.0.0.0:9092")
	go func() {
		log.Fatalln(gs.Serve(lis))
	}()
}
