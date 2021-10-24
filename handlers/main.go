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
	msg_err := "OK"
	url_file := os.Getenv("STATIONSURLS")
	station := req.GetStationName()
	url, err := readCsvFile(url_file, station)
	if err != nil {
		s.log.Error(url, "error", err)
		msg_err = err.Error()
		return &protobuff.GetStationWorkloadResponse{WorkLoad: map[string]*protobuff.DayWork{}, Error: url}, nil
	}
	res, err := GetMap(url, 2)
	if req.GetIsUpdateDB() == true {
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}
		defer conn.Close()
		tmpres, err := json.Marshal(res)
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}
		v, err := conn.Do("HSET", station, "WorkLoad", string(tmpres))
		fmt.Println(v)
		if err != nil {
			result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}

		fmt.Println("Added to DB")
	}
	if err != nil {
		s.log.Error(url, "error", err)
		msg_err = err.Error()
		return &protobuff.GetStationWorkloadResponse{WorkLoad: map[string]*protobuff.DayWork{}, Error: url}, nil
	}

	result := &protobuff.GetStationWorkloadResponse{WorkLoad: res, Error: msg_err}
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
