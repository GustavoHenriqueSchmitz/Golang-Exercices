package services

import (
	"api-cadastro-de-paises/database"
	"api-cadastro-de-paises/models"
	"errors"
)

func ReturnCities(jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	cities := []models.City{}
	err := database.DB.Table("api_countries.cities").Select("*").Scan(&cities)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ouve um erro ao carregar a lista de cidades."
		jsonreturn.Error = true
		return jsonreturn, errors.New("ERRO ao acessar/pegar dados do banco de dados.")
	}

	jsonreturn.Data = cities
	jsonreturn.Message = nil
	jsonreturn.Error = true

	return jsonreturn, nil
}

func ReturnOneCity(city_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	city := models.City{}
	err := database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar a cidade selecionada."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if city.Id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	city = models.City{}
	err = database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar a cidade selecionada."
		jsonreturn.Error = true
		return jsonreturn, errors.New("ERRO ao acessar/pegar dados do banco de dados.")
	}

	jsonreturn.Data = city
	jsonreturn.Message = nil
	jsonreturn.Error = false

	return jsonreturn, nil
}

func AddCity(city_add models.City, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	err := database.DB.Table("api_countries.cities").Select("name", "state_id").Create(&city_add)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro, ao criar nova cidade."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Cidade criada com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}

func DeleteCity(city_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	city := models.City{}
	err := database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar a cidade selecionada."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if city.Id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	city = models.City{}
	err = database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)
	if city.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de cidade específicado, não existe.")
	}

	city = models.City{}
	err = database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Delete(&city)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar a cidade selecionada."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Sucesso!"
	jsonreturn.Error = false
	return jsonreturn, nil

	return jsonreturn, nil
}

func UpdateCity(city_id int, city_update models.City, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	city := models.City{}
	err := database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar a cidade selecionada."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if city.Id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	city = models.City{}
	err = database.DB.Table("api_countries.cities").Select("*").Where("id", city_id).Scan(&city)
	if city.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de cidade específicado, não existe.")
	}

	err = database.DB.Table("api_countries.cities").Select("name", "state_id").Where("id", city_id).UpdateColumns(city_update)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro ao atualizar cidade."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if city_id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! A cidade específicada, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Cidade atualizada com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}
