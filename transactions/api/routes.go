package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"

	"github.com/galamshar/microservices-wallet/transactions/handlers"
	"github.com/galamshar/microservices-wallet/transactions/internal/middlewares"
)

func routes(DB *gorm.DB, RDB *redis.Client) *fiber.App {
	app := fiber.New()

	handler := handlers.NewTransactionsHandlerService(DB, RDB)

	user := app.Group("/transaction")

	user.Use(cors.New())

	user.Get("/:id/:page", middlewares.JWTMiddleware(), handler.GetTransactions) //Get Transactions
	user.Post("/", middlewares.JWTMiddleware(), handler.CreateTransaction)       //Create New transactions

	return app
}
