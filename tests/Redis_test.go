package tests

import (
	redisservice "RailwayStationsWorkload_micro/redis_service"
	redisProtobuff "RailwayStationsWorkload_micro/redis_service/protobuff"
	"context"
	"io"
	"testing"

	"google.golang.org/grpc"
)

// should be used only with no data in Redis to pass this test without accidential matches for test requests
// so using ClearRedis() to achieve that
func TestMissingData(t *testing.T) {
	redisservice.ClearRedis()
	redisservice.NewRedis()
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

	conn, err := grpc.Dial(redisservice.Redisport, grpc.WithBlock(), grpc.WithInsecure())
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

func TestRedisStoreAndSearch(t *testing.T) {
	redisservice.ClearRedis()
	redisservice.NewRedis()
	lengthCases := 3
	fillingCmds := []struct {
		request      redisProtobuff.StoreWorkloadRequest
		ErrorMessage string
	}{
		{
			request:      redisProtobuff.StoreWorkloadRequest{Station: "Дачное_11.11.2021", Workload: "something"},
			ErrorMessage: "OK",
		},
		{
			request:      redisProtobuff.StoreWorkloadRequest{Station: "Лигово_15.11.2021", Workload: "Other something"},
			ErrorMessage: "OK",
		},
		{
			request:      redisProtobuff.StoreWorkloadRequest{Station: "Дачное_15.11.2021", Workload: "Another something"},
			ErrorMessage: "OK",
		},
	}
	searchCmds := []struct {
		request  redisProtobuff.Stations
		response []redisProtobuff.SearchWorkloadResponse
	}{
		{
			request:  redisProtobuff.Stations{StationsNames: "Дачное"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Дачное_11.11.2021", Workload: "something"}},
		},
		{
			request:  redisProtobuff.Stations{StationsNames: "Лигово"},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Лигово_15.11.2021", Workload: "Other something"}},
		},
		{
			request:  redisProtobuff.Stations{StationsNames: "Дачное "},
			response: []redisProtobuff.SearchWorkloadResponse{{StationName: "Дачное_11.11.2021", Workload: "something"}, {StationName: "Дачное_15.11.2021", Workload: "Another something"}},
		},
	}
	conn, err := grpc.Dial(redisservice.Redisport, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		t.Error("Error when connecting to Redis service")
	}
	defer conn.Close()
	redis_client := redisProtobuff.NewRedisServiceClient(conn)
	//storing data
	for i := 0; i < lengthCases; i++ {
		respMsg, err := redis_client.StoreWorkload(context.Background(), &fillingCmds[i].request)
		if err != nil {
			t.Errorf("Internal Redis error when storing %v", &fillingCmds[i])
		}
		if respMsg.Error != fillingCmds[i].ErrorMessage {
			t.Errorf("Error %v returned from stroring %v", respMsg.Error, &fillingCmds[i].request)
		}
		//data stored now gonna search for it
		stream, err := redis_client.SearchWorkload(context.Background(), &searchCmds[i].request)
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
				t.Errorf("Error occures for test %v with error %v", searchCmds[i].request.StationsNames, err)
			}
			if feature.StationName != searchCmds[i].response[0].StationName {
				if feature.Workload != searchCmds[i].response[1].Workload {
					t.Errorf("Returned %v workload when it should be %v", feature.Workload, searchCmds[i].response[1].Workload)
				} else if feature.StationName != searchCmds[i].response[1].StationName {
					t.Errorf("%v when should be %v", feature.StationName, searchCmds[i].response[cnt].StationName)
				}

			} else if feature.Workload != searchCmds[i].response[0].Workload {
				t.Errorf("Returned %v workload when it should be %v", feature.Workload, searchCmds[i].response[0].Workload)
			}
			cnt++
		}
	}
}
