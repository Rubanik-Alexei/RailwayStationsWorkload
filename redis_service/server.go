package redisservice

import (
	"RailwayStationsWorkload_micro/redis_service/protobuff"
	"context"
	"fmt"

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

func (s *Server) StoreWorkload(ctx context.Context, req *protobuff.StoreWorkloadRequest) (*protobuff.ErrorMsg, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
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
