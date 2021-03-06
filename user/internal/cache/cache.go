package cache

import (
	"context"
	"log"

	"github.com/galamshar/microservices-wallet/user/internal/environment"

	"github.com/go-redis/redis/v8"
)

//NewRedisClient return a new redis client
func NewRedisClient() *redis.Client {
	addr := environment.AccessENV("REDIS_ADDR")

	log.Println(addr)

	if addr == "" {
		log.Fatalln("Error in Getting the ADDR from ENV")
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	status := rdb.Ping(context.Background())

	if status.Err() != nil {
		return nil
	}

	log.Println("Redis Cache Connected")

	return rdb
}
