package controllers

import (
	"api-cadastro-de-paises/models"
	"api-cadastro-de-paises/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ViewCountries(r *fiber.Ctx) error {
	jsonreturn, _ := services.ReturnAllCountries()
	r.JSON(jsonreturn)

	return nil
}

func ViewOneCountry(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("country_id")

	country_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar o país selecionado."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.ReturnOneCountrie(country_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func ViewStatesByCountry(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("country_id")

	country_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar os estados, do país selecionado."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.ReturnStatesByCountry(country_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func DeleteCountry(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("country_id")
	country_id, err := strconv.Atoi(url_id)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar o país selecionado."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.DeleteCountry(country_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func AddCountry(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	country_add := models.Country{}

	err := r.BodyParser(&country_add)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.AddCountry(country_add, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func UpdateCountry(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	country_update := models.Country{}
	url_id := r.Params("country_id")
	country_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao atualizar o país selecionado."
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	err = r.BodyParser(&country_update)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.UpdateCountry(country_id, country_update, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}
