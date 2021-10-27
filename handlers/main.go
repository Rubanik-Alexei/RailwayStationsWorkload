package handlers

import (
	"RailwayStationsWorkload/protobuff"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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
		result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: "Cannot connect to Redis"}}}
		return result, nil
	}
	defer conn.Close()
	if req.GetStationName() == "All" {
		keys, err := redis.Strings(conn.Do("KEYS", "*"))
		if err != nil || len(keys) == 0 {
			result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: "No data found in Redis"}}}
			return result, nil
		}
		tmpresp := []*protobuff.StationData{}
		for _, v := range keys {
			var tmp map[string]*protobuff.DayWork
			values, err := redis.String(conn.Do("HGET", v, "WorkLoad"))
			if err != nil {
				result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: v, RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()}}}
				return result, nil
			}
			//fmt.Println(values)
			//Unmarshaling to response struct
			err = json.Unmarshal([]byte(values), &tmp)
			if err != nil {
				result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: v, RespWorkLoad: map[string]*protobuff.DayWork{}, Error: fmt.Sprintf("Failed to retrieve workload for station %s", v)}}}
				return result, nil
			}
			tmpresp = append(tmpresp, &protobuff.StationData{RespstationName: v, RespWorkLoad: tmp, Error: "OK\n"})
		}
		result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: tmpresp}
		return result, nil
	}
	values, err := redis.String(conn.Do("HGET", req.GetStationName(), "WorkLoad"))
	if err != nil {
		result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: fmt.Sprintf("Failed to retrieve workload for station %s", req.GetStationName())}}}
		return result, nil
	}
	var tmp map[string]*protobuff.DayWork
	fmt.Println(values)
	//Unmarshaling to response struct
	err = json.Unmarshal([]byte(values), &tmp)
	if err != nil {
		result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()}}}
		return result, nil
	}
	result := &protobuff.GetStationWorkloadFromDBResponse{StationWorkloads: []*protobuff.StationData{{RespstationName: req.GetStationName(), RespWorkLoad: tmp, Error: msg_err}}}

	return result, nil
}

//using goroutines to retrieve and send data for each station
func (s *Server) GetManyStationWorkload(req *protobuff.GetStationWorkloadFromDBRequest, srv protobuff.MyService_GetManyStationWorkloadServer) error {
	stations_array := strings.Split(req.GetStationName(), ",")
	var wg sync.WaitGroup
	for _, v := range stations_array {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			url_file := os.Getenv("STATIONSURLS")
			url, err := ReadCsvFile(url_file, v)
			if err != nil {
				s.log.Error(url, "error", err)
				srv.Send(&protobuff.StationData{RespstationName: v, RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()})
			}
			res, err := GetMap(url, 2)
			if err != nil {
				s.log.Error(url, "error", err)
				srv.Send(&protobuff.StationData{RespstationName: v, RespWorkLoad: map[string]*protobuff.DayWork{}, Error: err.Error()})
			}
			resp := &protobuff.StationData{RespstationName: v, RespWorkLoad: res, Error: "OK"}
			if err := srv.Send(resp); err != nil {
				log.Printf("send error %v", err)
			}
			s.log.Info("Send workload for station : %s", v)
		}(v)
		//time.Sleep(5 * time.Second)
	}
	wg.Wait()
	return nil
}

func (s *Server) mustEmbedUnimplementedMyServiceServer() error {
	return nil
}
