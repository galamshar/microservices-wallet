package main

import (
	"github.com/galamshar/microservices-wallet/user/api"
	"github.com/galamshar/microservices-wallet/user/grpc"
)

func main() {
	go grpc.NewGRPCServer()
	api.Start()
}
