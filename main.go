package main

import (
	gatewayservice "RailwayStationsWorkload_micro/cmd/gateway"
	redisservice "RailwayStationsWorkload_micro/cmd/redis"
	workloadservice "RailwayStationsWorkload_micro/cmd/workload"
)

func main() {
	//starting all services
	workloadservice.NewWorkload()
	redisservice.NewRedis()
	gatewayservice.NewGateway()

}
