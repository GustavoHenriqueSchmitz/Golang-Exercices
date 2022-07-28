package routes

import (
	"api-cadastro-de-paises/controllers"

	"github.com/gofiber/fiber/v2"
)

func Requests() {

	router := fiber.New()

	CountriesRoutes(router.Group("/countries"))
	StatesRoutes(router.Group("/states"))
	CitiesRoutes(router.Group("/cities"))

	router.Get("/", controllers.Home)

	router.Listen(":8080")
}
