package main

import (
	gatewayservice "RailwayStationsWorkload_micro/gateway_service"
	redisservice "RailwayStationsWorkload_micro/redis_service"
	workloadservice "RailwayStationsWorkload_micro/workload_service"
)

func main() {
	//starting all services
	workloadservice.NewWorkload()
	redisservice.NewRedis()
	gatewayservice.NewGateway()

}
