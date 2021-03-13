package main

import (
	"log"

	"github.com/galamshar/microservices-wallet/auth/api"
	"github.com/galamshar/microservices-wallet/auth/grpc"
)

func main() {
	log.Println("Start Auth gRPC")
	go grpc.NewGRPCServer()
	log.Println("Start Auth Server")
	api.Start()
}
