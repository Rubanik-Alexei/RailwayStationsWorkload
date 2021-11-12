package redisservice

import (
	"RailwayStationsWorkload_micro/redis_service/protobuff"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	log hclog.Logger
	protobuff.UnimplementedRedisServiceServer
}

func NewMyServer(l hclog.Logger) *Server {
	return &Server{l, protobuff.UnimplementedRedisServiceServer{}}
}

type MyError struct {
	error_msg string
}

func (m MyError) Error() string {
	return m.error_msg
}

const redisAdress = "localhost:6379"

func (s *Server) StoreWorkload(ctx context.Context, req *protobuff.StoreWorkloadRequest) (*protobuff.ErrorMsg, error) {
	conn, err := redis.Dial("tcp", redisAdress)
	errMsg := "OK"
	if err != nil {
		errMsg = "Cannot connect to Redis"
		return &protobuff.ErrorMsg{Error: errMsg}, err
	}
	defer conn.Close()
	v, err := conn.Do("HSET", req.GetStation(), "WorkLoad", req.GetWorkload())
	fmt.Println(v)
	if err != nil {
		errMsg = "Cannot add data to Redis"
		return &protobuff.ErrorMsg{Error: errMsg}, err
	}
	return &protobuff.ErrorMsg{Error: errMsg}, nil
}

func (s *Server) SearchWorkload(req *protobuff.Stations, srv protobuff.RedisService_SearchWorkloadServer) error {

	stations_array := strings.Split(req.GetStationsNames(), ",")
	var wg sync.WaitGroup
	for _, v := range stations_array {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			conn, err := redis.Dial("tcp", redisAdress)
			if err != nil {
				log.Printf("Cannot connect to redis")
			}
			defer conn.Close()
			wl, err := redis.String(conn.Do("HGET", v, "WorkLoad"))
			if err != nil {
				log.Printf("Got error %v when searching for %v", err, v)
			}
			msg := &protobuff.SearchWorkloadResponse{StationName: v, Workload: wl}
			if err := srv.Send(msg); err != nil {
				log.Printf("send error %v", err)
			}
			s.log.Info("Send workload for station : %s", v)
		}(v)
	}
	wg.Wait()
	return nil
}
