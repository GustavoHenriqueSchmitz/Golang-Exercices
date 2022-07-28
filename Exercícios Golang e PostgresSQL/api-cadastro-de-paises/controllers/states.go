package controllers

import (
	"api-cadastro-de-paises/models"
	"api-cadastro-de-paises/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ViewStates(r *fiber.Ctx) error {
	jsonreturn, _ := services.ReturnAllStates()
	r.JSON(jsonreturn)

	return nil
}

func ViewOneState(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("state_id")

	state_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar o estado selecionado."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.ReturnOneState(state_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func DeleteState(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	url_id := r.Params("state_id")

	state_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar o estado selecionado."
		jsonreturn.Error = true
		r.JSON(jsonreturn)

		return nil
	}

	jsonreturn, _ = services.DeleteState(state_id, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func AddState(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	state_add := models.States{}

	err := r.BodyParser(&state_add)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.AddState(state_add, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}

func UpdateState(r *fiber.Ctx) error {

	jsonreturn := models.Jsonreturn{}
	state_update := models.States{}
	url_id := r.Params("state_id")

	state_id, err := strconv.Atoi(url_id)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao atualizar o estado selecionado."
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	err = r.BodyParser(&state_update)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.UpdateState(state_id, state_update, jsonreturn)
	r.JSON(jsonreturn)

	return nil
}
