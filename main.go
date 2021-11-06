package main

import (
	gatewayservice "RailwayStationsWorkload_micro/gateway_service"
	workloadservice "RailwayStationsWorkload_micro/workload_service"
)

func main() {
	//starting all services
	workloadservice.NewWorkload()
	gatewayservice.NewGateway()
}
