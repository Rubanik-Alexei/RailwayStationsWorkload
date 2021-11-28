package tests

import (
	redisworkload "RailwayStationsWorkload_micro/redis_service"
	redisProtobuff "RailwayStationsWorkload_micro/redis_service/protobuff"
	"context"
	"io"
	"testing"

	"google.golang.org/grpc"
)

// should be used only with no data in Redis to pass this test without accidential matches for test requests
func TestMissingData(t *testing.T) {
	redisworkload.NewRedis()
	tests := []struct {
		request  redisProtobuff.Stations
		response []redisProtobuff.SearchWorkloadResponse
	}{
		{
			request:  redisProtobuff.Stations{StationsNames: "Спб"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Спб", Workload: ""}},
		},
		//Should split only by ","
		{
			request:  redisProtobuff.Stations{StationsNames: "Дачное Лигово"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Дачное Лигово", Workload: ""}},
		},
		{
			request:  redisProtobuff.Stations{StationsNames: "Дачное, Лигово"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Дачное", Workload: ""}, {StationName: "Лигово", Workload: ""}},
		},
		{
			request:  redisProtobuff.Stations{StationsNames: "Дачное,Лигово"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Дачное", Workload: ""}, {StationName: "Лигово", Workload: ""}},
		},
	}

	conn, err := grpc.Dial(redisworkload.Redisport, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		t.Error("Error when connecting to Redis service")
	}
	defer conn.Close()
	redis_client := redisProtobuff.NewRedisServiceClient(conn)
	for _, tt := range tests {
		stream, err := redis_client.SearchWorkload(context.Background(), &tt.request)
		if err != nil {
			t.Error("Bad response from Redis service")
		}
		cnt := 0
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("Error occures for test %v with error %v", tt.request.StationsNames, err)
			}
			if feature.StationName != tt.response[0].StationName {
				if feature.StationName != tt.response[1].StationName {
					t.Errorf("%v when should be %v", feature.StationName, tt.response[cnt].StationName)
				}
			}
			if feature.Workload != tt.response[cnt].Workload {
				t.Error("Nonempty workload when should be empty")
			}
			cnt++
		}

	}
}
