package server

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"

	"github.com/galamshar/microservices-wallet/auth/internal/cache"
	internal "github.com/galamshar/microservices-wallet/auth/internal/storage"
	"github.com/galamshar/microservices-wallet/auth/storage"
	"github.com/jinzhu/gorm"
)

//Server User Server struct
type Server struct {
}

var storageService *storage.AuthStorageService

//GetStorageService Start the storage service for GPRC server
func GetStorageService() {
	var DB *gorm.DB = internal.ConnectDB()
	var RDB *redis.Client = cache.NewRedisClient()

	storageService = storage.NewAuthStorageService(DB, RDB)
}

//CloseDB Close both DB
func CloseDB() {
	storageService.CloseDB()
}

//ChangeAuthCache Change in redis the User's Username or email
func (s *Server) ChangeAuthCache(ctx context.Context, request *NewUserInfo) (*AuthResponse, error) {
	if len(request.OldUsername) > 0 && len(request.NewUsername) > 0 {
		success, err := storageService.ChangeRegisterCache(request.OldUsername, request.NewUsername, "", "")
		return &AuthResponse{Success: success}, err
	}

	if len(request.OldEmail) > 0 && len(request.NewEmail) > 0 {
		success, err := storageService.ChangeRegisterCache("", "", request.NewEmail, request.OldEmail)
		return &AuthResponse{Success: success}, err
	}

	return &AuthResponse{Success: false}, errors.New("Invalid Input")
}
