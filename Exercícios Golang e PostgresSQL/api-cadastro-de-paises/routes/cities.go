package routes

import (
	"api-cadastro-de-paises/controllers"

	"github.com/gofiber/fiber/v2"
)

func CitiesRoutes(r fiber.Router) {
	r.Get("/", controllers.ViewCities)
	r.Get("/:city_id", controllers.ViewOneCity)
	r.Put("/add", controllers.AddCity)
	r.Delete("/delete/:city_id", controllers.DeleteCity)
	r.Put("/update/:city_id", controllers.UpdateCity)
}
