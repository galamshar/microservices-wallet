package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"

	"github.com/galamshar/microservices-wallet/auth/internal/cache"
	"github.com/galamshar/microservices-wallet/auth/internal/storage"
)

//Start Start a new User server API
func Start() {
	var DB *gorm.DB

	DB = storage.ConnectDB()
	defer DB.Close()

	var RDB *redis.Client

	RDB = cache.NewRedisClient()
	defer RDB.Close()

	app := routes(DB, RDB)
	createServer(app)
}
