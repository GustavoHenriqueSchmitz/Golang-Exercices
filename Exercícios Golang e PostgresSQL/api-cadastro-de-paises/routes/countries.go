package routes

import (
	"api-cadastro-de-paises/controllers"

	"github.com/gofiber/fiber/v2"
)

func CountriesRoutes(r fiber.Router) {
	r.Get("/", controllers.ViewCountries)
	r.Get("/:country_id", controllers.ViewOneCountry)
	r.Get("/states/:country_id", controllers.ViewStatesByCountry)
	r.Put("/add", controllers.AddCountry)
	r.Delete("/delete/:country_id", controllers.DeleteCountry)
	r.Put("/update/:country_id", controllers.UpdateCountry)
}
