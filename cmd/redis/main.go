package redisservice

import (
	"RailwayStationsWorkload_micro/config"
	redisservice "RailwayStationsWorkload_micro/redis_service"
	"RailwayStationsWorkload_micro/redis_service/protobuff"
	"log"
	"net"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func NewRedis() {
	logg := hclog.Default()
	lis, err := net.Listen("tcp", config.Redisport)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	gs := grpc.NewServer()
	ms := redisservice.NewMyServer(logg)
	protobuff.RegisterRedisServiceServer(gs, ms)
	//reflection.Register(reflection.GRPCServer(gs))
	// Serve gRPC server
	log.Println("Serving Redis_gRPC on 0.0.0.0:9091")
	go func() {
		log.Fatalln(gs.Serve(lis))
	}()
}
