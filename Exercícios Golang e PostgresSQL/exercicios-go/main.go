package main

import (
	"log"

	// _ "exercicios/docs" // load API Docs files (Swagger)

	"github.com/gofiber/fiber/v2"
	"gitlab.com/egadnet/egadnet-go-core/pkg/configs"
	"gitlab.com/egadnet/egadnet-go-core/pkg/middleware"
	coreRoutes "gitlab.com/egadnet/egadnet-go-core/pkg/routes"
	"gitlab.com/egadnet/egadnet-go-core/pkg/utils"
	redisutils "gitlab.com/egadnet/egadnet-go-core/pkg/utils/redis"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title EgadEduc GO Courses Admin
// @version 1.0
// @description The Courses Admin API.
// @termsOfService http://swagger.io/terms/
// @contact.name Mauricio Schmitz
// @contact.email mauricio@egadnet.com.br
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	// Redis
	_, err := redisutils.OpenRedisConnection()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	// Routes.
	coreRoutes.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	// routes.Inicio(app)           // Register a courses routes for app.

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
