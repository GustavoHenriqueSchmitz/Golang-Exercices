package controllers

import (
	"api-cadastro-de-paises/models"
	"api-cadastro-de-paises/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ViewCities(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	jsonreturn, _ = services.ReturnCities(jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func ViewOneCity(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("city_id")

	city_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar a cidade selecionada."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.ReturnOneCity(city_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func AddCity(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	city_add := models.City{}

	err := r.BodyParser(&city_add)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.AddCity(city_add, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func DeleteCity(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("city_id")
	city_id, err := strconv.Atoi(url_id)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar a cidade selecionada."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.DeleteCity(city_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func UpdateCity(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	city_update := models.City{}
	url_id := r.Params("city_id")
	city_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao atualizar a cidade selecionada."
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	err = r.BodyParser(&city_update)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.UpdateCity(city_id, city_update, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}
