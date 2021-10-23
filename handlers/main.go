package handlers

import (
	"RailwayStationsWorkload/protobuf"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	log hclog.Logger
}

func NewMyServer(l hclog.Logger) *Server {
	return &Server{l}
}

// func main() {
//     records := readCsvFile("../tasks.csv")
//     fmt.Println(records)
// }
func (s *Server) GetStationWorkload(ctx context.Context, req *protobuf.GetStationWorkloadRequest) (*protobuf.GetStationWorkloadResponse, error) {
	msg_err := "OK"
	url_file := os.Getenv("STATIONSURLS")
	station := req.GetStationName()
	url, err := readCsvFile(url_file, station)
	if err != nil {
		s.log.Error(url, "error", err)
		msg_err = err.Error()
		return &protobuf.GetStationWorkloadResponse{WorkLoad: map[string]*protobuf.DayWork{}, Error: url}, nil
	}
	res, err := GetMap(url, 2)
	if req.GetIsUpdateDB() == true {
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			result := &protobuf.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}
		// Importantly, use defer to ensure the connection is always
		// properly closed before exiting the main() function.
		defer conn.Close()

		// Send our command across the connection. The first parameter to
		// Do() is always the name of the Redis command (in this example
		// HMSET), optionally followed by any necessary arguments (in this
		// example the key, followed by the various hash fields and values).
		tmpres, err := json.Marshal(res)
		if err != nil {
			result := &protobuf.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}
		v, err := conn.Do("HSET", station, "WorkLoad", string(tmpres))
		fmt.Println(v)
		if err != nil {
			result := &protobuf.GetStationWorkloadResponse{WorkLoad: res, Error: err.Error()}
			return result, nil
		}

		fmt.Println("Все ок")
	}
	if err != nil {
		s.log.Error(url, "error", err)
		msg_err = err.Error()
		return &protobuf.GetStationWorkloadResponse{WorkLoad: map[string]*protobuf.DayWork{}, Error: url}, nil
	}
	result := &protobuf.GetStationWorkloadResponse{WorkLoad: res, Error: msg_err}
	return result, nil

}
func (s *Server) GetStationWorkloadFromDB(ctx context.Context, req *protobuf.GetStationWorkloadFromDBRequest) (*protobuf.GetStationWorkloadFromDBResponse, error) {
	msg_err := "OK"
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		result := &protobuf.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuf.DayWork{}, Error: err.Error()}
		return result, nil
	}
	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).
	values, err := redis.String(conn.Do("HGET", req.GetStationName(), "WorkLoad"))

	if err != nil {
		result := &protobuf.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuf.DayWork{}, Error: err.Error()}
		return result, nil
	}
	var tmp map[string]*protobuf.DayWork
	fmt.Println(values)
	err = json.Unmarshal([]byte(values), &tmp)
	if err != nil {
		result := &protobuf.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuf.DayWork{}, Error: err.Error()}
		return result, nil
	}
	//
	if err != nil {
		result := &protobuf.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: map[string]*protobuf.DayWork{}, Error: err.Error()}
		return result, nil
	}
	result := &protobuf.GetStationWorkloadFromDBResponse{RespstationName: req.GetStationName(), RespWorkLoad: tmp, Error: msg_err}
	return result, nil
}

// func (s *Server) Get(ctx context.Context, req *protobuf.GetStationAdressesRequest) (*protobuf.GetStationAdressesResponse, error) {
// 	s.log.Info("Handle GetAdresses", "for station", req.GetStationname())
// 	res := &protobuf.Item{Adress: req.GetStationname()}
// 	return &protobuf.GetStationAdressesResponse{Items: []*protobuf.Item{res}}, nil

// }

// func (s *Server) GetStationsAmount(ctx context.Context, req *protobuf.GetStationsAmountRequest) (*protobuf.GetStationsAmountResponse, error) {
// 	s.log.Info("Handle GetAmount", "for station", req.GetId())
// 	res := &protobuf.Amount{Station: "Пунк", Amount: req.GetId()}
// 	return &protobuf.GetStationsAmountResponse{Amounts: []*protobuf.Amount{res}}, nil
// }
