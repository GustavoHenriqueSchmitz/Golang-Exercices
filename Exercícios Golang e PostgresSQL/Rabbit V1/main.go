package main

import (
	"rabbitmq/configs"
	"rabbitmq/middleware"
	"rabbitmq/routes"
	"rabbitmq/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.FiberConfig()
	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	routes.CreateProofs(app)
	utils.StartServer(app)
}
