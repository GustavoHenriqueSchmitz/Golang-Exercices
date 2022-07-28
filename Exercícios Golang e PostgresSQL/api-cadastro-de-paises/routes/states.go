package routes

import (
	"api-cadastro-de-paises/controllers"

	"github.com/gofiber/fiber/v2"
)

func StatesRoutes(r fiber.Router) {
	r.Get("/", controllers.ViewStates)
	r.Get("/:state_id", controllers.ViewOneState)
	r.Put("/add", controllers.AddState)
	r.Delete("/delete/:state_id", controllers.DeleteState)
	r.Put("/update/:state_id", controllers.UpdateState)
}
