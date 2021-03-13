package client

import (
	"context"
	"log"

	"github.com/galamshar/microservices-wallet/user/internal/environment"
	"github.com/galamshar/microservices-wallet/user/models"
	"google.golang.org/grpc"
)

var moveClient models.MovementServiceClient
var moveConn *grpc.ClientConn

//StartMoveClient Start the client for movement gRPC
func startMoveClient() {
	urlTarget := environment.AccessENV("MOVEMENT_GRPC")

	if urlTarget == "" {
		log.Fatalln("Error in Access to GRPC URL in User Service")
	}
	moveConn, err := grpc.Dial(urlTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	moveClient = models.NewMovementServiceClient(moveConn)
}

//CloseMoveClient Close the client for movement gRPC
func closeMoveClient() {
	moveConn.Close()
}

//CreateMovement Create a new movement in DB
func CreateMovement(relation string, change string, origin string) (bool, error) {
	newMovement := &models.MovementRequest{
		Relation: relation,
		Change:   change,
		Origin:   origin,
	}

	response, err := moveClient.CreateMovement(context.Background(), newMovement)

	if err != nil {
		return false, err
	}

	return response.Sucess, nil
}
