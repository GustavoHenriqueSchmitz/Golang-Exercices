package services

import (
	"api-cadastro-de-paises/database"
	"api-cadastro-de-paises/models"
	"errors"
)

func ReturnAllStates() (models.Jsonreturn, error) {

	states := []models.States{}
	jsonreturn := models.Jsonreturn{}

	err := database.DB.Table("api_countries.states").Select("*").Scan(&states)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ouve um erro ao carregar a lista de estados."
		jsonreturn.Error = true
		return jsonreturn, errors.New("ERRO ao acessar/pegar dados do banco de dados.")
	}

	jsonreturn.Data = states
	jsonreturn.Message = nil
	jsonreturn.Error = false

	return jsonreturn, nil
}

func ReturnOneState(state_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	state := models.States{}
	err := database.DB.Table("api_countries.states").Select("*").Where("id", state_id).Scan(&state)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar o estado selecionado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if state.Id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O estado específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = state
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func DeleteState(state_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	state := models.States{}
	err := database.DB.Table("api_countries.states").Select("*").Where("id", state_id).Scan(&state)
	if state.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O estado específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de estado específicado, não existe.")
	}

	states := models.States{}
	err = database.DB.Table("api_countries.states").Select("*").Where("id", state_id).Delete(&states)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar o estado selecionado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Sucesso!"
	jsonreturn.Error = false
	return jsonreturn, nil
}

func UpdateState(state_id int, state_update models.States, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	state := models.States{}
	err := database.DB.Table("api_countries.states").Select("*").Where("id", state_id).Scan(&state)
	if state.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O estado específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de estado específicado, não existe.")
	}

	err = database.DB.Table("api_countries.states").Select("name", "acronym", "country_id").Where("id", state_id).UpdateColumns(state_update)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro, ao atualizar estado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if state_id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O estado específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Estado atualizado com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}

func AddState(state models.States, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	err := database.DB.Table("api_countries.states").Select("name", "acronym", "country_id").Create(&state)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro, ao criar novo estado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Estado criado com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}
