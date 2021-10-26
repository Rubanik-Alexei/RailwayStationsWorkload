package handlers

import (
	"RailwayStationsWorkload/protobuff"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	log hclog.Logger
	protobuff.UnimplementedMyServiceServer
}

func NewMyServer(l hclog.Logger) *Server {
	return &Server{l, protobuff.UnimplementedMyServiceServer{}}
}

func (s *Server) GetStationWorkload(ctx context.Context, req *protobuff.GetStationWorkloadRequest) (*protobuff.GetStationWorkloadResponse, error) {
	//retrieving path to file with needed stations and their urls
	url_file := os.Getenv("STATIONSURLS")
	station := req.GetStationName()
	url, err := ReadCsvFile(url_file, station)
	if err != nil {
		s.log.Error(url, "error", err)
		return &protobuff.GetStationWorkloadResponse{WorkLoad: map[string]*protobuff.DayWork{}, Error: url}, nil
	}
	res, err := GetMap(url, 2)
	//Adding data to Redis if requested
	if req.GetIsUpdateDB() {
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: "Cannot connect to Redis"}
			return result, nil
		}
		defer conn.Close()
		//Needed to marshal response struct to json to be able to store it
		tmpres, err := json.Marshal(res)
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: "Cannot parse response"}
			return result, nil
		}
		v, err := conn.Do("HSET", station, "WorkLoad", string(tmpres))
		fmt.Println(v)
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: "Cannot add data to Redis"}
			return result, nil
		}

		//fmt.Println("Added to DB")
	}
	if err != nil {
		s.log.Error(url, "error", err)
		return &protobuff.GetStationWorkloadResponse{WorkLoad: map[string]*protobuff.DayWork{}, Error: url}, nil
	}
	result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: "OK"}
	//Marshaling for browser availability
	_, _ = protojson.Marshal(result)
	// fmt.Println(string(jsontmp))
	return result, nil

}
func (s *Server) GetStationWorkloadFromDB(ctx context.Context, req *protobuff.GetStationWorkloadFromDBRequest) (*protobuff.GetStationWorkloadFromDBResponse, error) {
	msg_err := "OK"
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		result := &protobuff.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()}
		return result, nil
	}
	defer conn.Close()
	values, err := redis.String(conn.Do("HGET", req.GetStationName(), "WorkLoad"))
	if err != nil {
		result := &protobuff.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()}
		return result, nil
	}
	var tmp map[string]*protobuff.DayWork
	fmt.Println(values)
	//Unmarshaling to response struct
	err = json.Unmarshal([]byte(values), &tmp)
	if err != nil {
		result := &protobuff.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()}
		return result, nil
	}
	result := &protobuff.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: tmp, Error: msg_err}

	return result, nil

}
func (s *Server) mustEmbedUnimplementedMyServiceServer() error {
	return nil
}
