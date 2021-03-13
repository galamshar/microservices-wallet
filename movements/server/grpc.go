package server

import (
	"log"
	"net"

	MoveServer "github.com/galamshar/microservices-wallet/movements/grpc"
	"google.golang.org/grpc"
)

//NewGRPCServer Create new gRPC server
func NewGRPCServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on Port :9000 %v", err)
	}

	s := MoveServer.Server{}

	grpcServer := grpc.NewServer()

	MoveServer.StartDB()
	defer MoveServer.CloseDB()

	MoveServer.RegisterMovementServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on Port :9000 %v", err)
	}
}
