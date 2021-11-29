package redisservice

import (
	"RailwayStationsWorkload_micro/redis_service/protobuff"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

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

//Clearing Redis (only for testing purpose)
func ClearRedis() {
	conn, err := redis.Dial("tcp", redisAdress)
	if err != nil {
		log.Fatal("Error when connecting to Redis service")
	}
	defer conn.Close()
	_, err = conn.Do("FLUSHALL")
	if err != nil {
		log.Fatal("Error when clearing database")
	}
}

//helper function for searching
func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func (s *Server) StoreWorkload(ctx context.Context, req *protobuff.StoreWorkloadRequest) (*protobuff.ErrorMsg, error) {
	conn, err := redis.Dial("tcp", redisAdress)
	errMsg := "OK"
	if err != nil {
		errMsg = "Cannot connect to Redis"
		return &protobuff.ErrorMsg{Error: errMsg}, err
	}
	defer conn.Close()

	//version of storing that is using for testing
	//v, err := conn.Do("SET", strings.TrimSpace(req.GetStation()), req.GetWorkload())

	v, err := conn.Do("SET", req.GetStation()+time.Now().Format("01-02-2006"), req.GetWorkload())

	fmt.Println(v)
	if err != nil {
		errMsg = "Cannot add data to Redis"
		return &protobuff.ErrorMsg{Error: errMsg}, err
	}
	return &protobuff.ErrorMsg{Error: errMsg}, nil
}

func (s *Server) SearchWorkload(req *protobuff.Stations, srv protobuff.RedisService_SearchWorkloadServer) error {

	stations_array := strings.Split(req.GetStationsNames(), ",")
	for i := range stations_array {
		stations_array[i] = strings.TrimSpace(stations_array[i])
	}
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
			keys, err := redis.Values(conn.Do("SCAN", 0, "MATCH", v+"*"))
			if err != nil {
				log.Printf("Error HSCAN %v", err)
			}
			for i, key := range keys {
				if i == 0 {
					continue
				}
				typedkey := key.([]interface{})
				tmp := typedkey[0].([]uint8)
				_ = tmp
				strkey := B2S(tmp)
				fmt.Println(strkey)
				wl, err := redis.String(conn.Do("GET", strkey))
				if err != nil {
					log.Printf("Got error %v when searching for %v", err, v)
				}
				msg := &protobuff.SearchWorkloadResponse{StationName: strkey, Workload: wl}
				if err := srv.Send(msg); err != nil {
					log.Printf("send error %v", err)
				}
				s.log.Info("Send workload for station : %s", strkey)

			}

		}(v)
	}
	wg.Wait()
	return nil
}
