package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(r *fiber.Ctx) error {
	return r.SendString("Página Inicial")
}
