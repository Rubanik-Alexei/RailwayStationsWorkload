package main

import (
	"RailwayStationsWorkload/handlers"
	"RailwayStationsWorkload/protobuff"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
)

// func main() {
// 	log := hclog.Default()
// 	gs := grpc.NewServer()
// 	ms := handlers.NewMyServer(log)

// 	protobuff.RegisterMyServiceServer(gs, ms)

// 	reflection.Register(gs)

// 	l, err := net.Listen("tcp", ":9092")
// 	if err != nil {
// 		log.Error("Unable to listen", "error", err)
// 		os.Exit(1)
// 	}
// 	gs.Serve(l)

// }

func main() {
	logg := hclog.Default()
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	gs := grpc.NewServer()
	ms := handlers.NewMyServer(logg)
	protobuff.RegisterMyServiceServer(gs, ms)
	//reflection.Register(reflection.GRPCServer(gs))
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9092")
	go func() {
		log.Fatalln(gs.Serve(lis))
	}()

	// Create a client connection to the gRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9092",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register handler
	err = protobuff.RegisterMyServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
