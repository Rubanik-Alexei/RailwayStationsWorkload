package main

import (
	"RailwayStationsWorkload/protobuff"
	"log"

	"context"
	//"google.golang.org/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main_() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := protobuff.NewMyServiceClient(conn)

	response, err := client.GetStationWorkloadFromDB(context.Background(), &protobuff.GetStationWorkloadFromDBRequest{StationName: "Славянка"})
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}
	tmp, _ := protojson.Marshal(response)
	// fmt.Println(string(jsontmp))
	log.Printf("Response from server: %s", string(tmp))

}
