package routes

import (
	"api-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func Requests() {

	router := fiber.New()

	router.Get("/", controllers.Home)
	router.Get("/ex01", controllers.Ex01)
	router.Get("/ex02", controllers.Ex02)
	router.Get("/ex03", controllers.Ex03)
	router.Get("/ex04", controllers.Ex04)
	router.Get("/ex05", controllers.Ex05)
	router.Get("/ex06", controllers.Ex06)
	router.Get("/ex07", controllers.Ex07)
	router.Get("/ex08", controllers.Ex08)
	router.Get("/ex09", controllers.Ex09)
	router.Get("/ex10", controllers.Ex10)
	router.Get("/ex11", controllers.Ex11)
	router.Get("/ex12", controllers.Ex12)
	router.Get("/ex13", controllers.Ex13)
	router.Get("/ex14", controllers.Ex14)
	router.Get("/ex15", controllers.Ex15)

	router.Listen(":8080")
}
