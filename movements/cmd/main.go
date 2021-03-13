package main

import (
	"log"

	"github.com/galamshar/microservices-wallet/movements/server"
)

func main() {
	log.Println("Start Movement gRPC Server")

	server.NewGRPCServer()
}
