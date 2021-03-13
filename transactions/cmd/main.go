package main

import (
	"github.com/galamshar/microservices-wallet/transactions/api"
	"github.com/galamshar/microservices-wallet/transactions/grpc"
)

func main() {
	go grpc.NewGRPCServer()
	api.Start()
}
