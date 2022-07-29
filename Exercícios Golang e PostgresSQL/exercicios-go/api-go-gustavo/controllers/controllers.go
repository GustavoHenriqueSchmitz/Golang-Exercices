package controllers

import (
	"api-go-gustavo/models"
	"api-go-gustavo/services"
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	jsonreturn := models.Jsonreturn{}
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	jsonreturn, _ = services.ValidaIdade(jsonreturn, user)
	json.NewEncoder(w).Encode(jsonreturn)
}

func Ex01(w http.ResponseWriter, r *http.Request) {
	jsonreturn := models.Jsonreturn{}
	nums := models.Nums{}
	json.NewDecoder(r.Body).Decode(&nums)
	jsonreturn, _ = services.SomaNumeros(jsonreturn, nums)
	json.NewEncoder(w).Encode(jsonreturn)
}

func Ex02(w http.ResponseWriter, r *http.Request) {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}
	json.NewDecoder(r.Body).Decode(&num)
	jsonreturn, _ = services.VerificaDivisao(jsonreturn, num)
	json.NewEncoder(w).Encode(jsonreturn)
}

func Ex03(w http.ResponseWriter, r *http.Request) {
	jsonreturn := models.Jsonreturn{}
	pessoa := models.Pessoa{}
	json.NewDecoder(r.Body).Decode(&pessoa)
	jsonreturn, _ = services.ValidaPessoa(jsonreturn, pessoa)
	json.NewEncoder(w).Encode(jsonreturn)
}

func Ex04(w http.ResponseWriter, r *http.Request) {
}

func Ex05(w http.ResponseWriter, r *http.Request) {

}

func Ex06(w http.ResponseWriter, r *http.Request) {

}

func Ex07(w http.ResponseWriter, r *http.Request) {

}

func Ex08(w http.ResponseWriter, r *http.Request) {

}

func Ex09(w http.ResponseWriter, r *http.Request) {

}

func Ex10(w http.ResponseWriter, r *http.Request) {

}

func Ex11(w http.ResponseWriter, r *http.Request) {

}

func Ex12(w http.ResponseWriter, r *http.Request) {

}

func Ex13(w http.ResponseWriter, r *http.Request) {

}

func Ex14(w http.ResponseWriter, r *http.Request) {

}

func Ex15(w http.ResponseWriter, r *http.Request) {

}
