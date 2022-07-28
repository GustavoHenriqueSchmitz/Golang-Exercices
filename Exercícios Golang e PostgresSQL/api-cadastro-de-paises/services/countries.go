package services

import (
	"api-cadastro-de-paises/database"
	"api-cadastro-de-paises/models"
	"errors"
)

func ReturnAllCountries() (models.Jsonreturn, error) {

	countries := []models.Country{}
	jsonreturn := models.Jsonreturn{}

	err := database.DB.Table("api_countries.countries").Select("*").Scan(&countries)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ouve um erro ao carregar a lista de países."
		jsonreturn.Error = true
		return jsonreturn, errors.New("ERRO ao acessar/pegar dados do banco de dados.")
	}

	jsonreturn.Data = countries
	jsonreturn.Message = nil
	jsonreturn.Error = false

	return jsonreturn, nil
}

func ReturnOneCountrie(country_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	country := models.Country{}
	err := database.DB.Table("api_countries.countries").Select("*").Where("id", country_id).Scan(&country)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao visualizar o país selecionado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if country.Id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O país específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = country
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func ReturnStatesByCountry(country_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	country := models.Country{}
	err := database.DB.Table("api_countries.countries").Select("*").Where("id", country_id).Scan(&country)
	if country.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O país específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de país específicado, não existe.")
	}

	states := []models.States{}

	err = database.DB.Table("api_countries.states").Select("id", "name", "acronym", "country_id").Where("country_id", country_id).Scan(&states)

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

func DeleteCountry(country_id int, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	country := models.Country{}
	err := database.DB.Table("api_countries.countries").Select("*").Where("id", country_id).Scan(&country)
	if country.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O país específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de país específicado, não existe.")
	}

	countries := models.Country{}
	err = database.DB.Table("api_countries.countries").Select("*").Where("id", country_id).Delete(&countries)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Ouve um erro ao apagar o país selecionado."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Sucesso!"
	jsonreturn.Error = false
	return jsonreturn, nil
}

func UpdateCountry(country_id int, country_update models.Country, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	country := models.Country{}
	err := database.DB.Table("api_countries.countries").Select("*").Where("id", country_id).Scan(&country)
	if country.Id <= 0 || err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O país específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! O id de país específicado, não existe.")
	}

	err = database.DB.Table("api_countries.countries").Select("name", "acronym").Where("id", country_id).UpdateColumns(country_update)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro, ao atualizar país."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	} else if country_id <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! O país específicado, não existe."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "País atualizado com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}

func AddCountry(country models.Country, jsonreturn models.Jsonreturn) (models.Jsonreturn, error) {

	err := database.DB.Table("api_countries.countries").Select("name", "acronym").Create(&country)

	if err.Error != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Erro, ao criar novo país."
		jsonreturn.Error = true
		return jsonreturn, err.Error
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "País criado com sucesso!"
	jsonreturn.Error = false

	return jsonreturn, nil
}
