package services

import (
	"api-go-gustavo/models"
	"errors"
	"strings"
)

func ValidaIdade(jsonreturn models.Jsonreturn, user models.User) (models.Jsonreturn, error) {

	if user.Name == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Nome Indefinido"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Nome indefinido")
	} else if user.Age <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Idade Ínvalida."
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Idade Ínvalida")
	} else if user.Age < 18 {
		user.Permissao = false
		jsonreturn.Data = user
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else if user.Age >= 18 {
		user.Permissao = true
		jsonreturn.Data = user
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	}

	jsonreturn.Data = nil
	jsonreturn.Message = "Ouve um erro ao validar"
	jsonreturn.Error = true

	return jsonreturn, errors.New("Ouve um erro ao validar.")
}

func SomaNumeros(jsonreturn models.Jsonreturn, nums models.Nums) (models.Jsonreturn, error) {

	soma := nums.N1 + nums.N2

	if soma > 20 {
		jsonreturn.Data = soma + 8
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else if soma <= 20 {
		jsonreturn.Data = soma - 5
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	}

	return jsonreturn, errors.New("Ouve um erro ao somar.")
}

func VerificaDivisao(jsonreturn models.Jsonreturn, num models.Nums) (models.Jsonreturn, error) {
	ndiv := []int{}

	if num.N1 <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "0 ou valores abaixo, não podem ser divididos"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Valor 0 ou abaixo, não pode ser dividido")
	}

	if num.N1%2 == 0 {
		ndiv = append(ndiv, 2)
	}

	if num.N1%5 == 0 {
		ndiv = append(ndiv, 5)
	}

	if num.N1%10 == 0 {
		ndiv = append(ndiv, 10)
	}

	if len(ndiv) == 0 {
		jsonreturn.Data = ndiv
		jsonreturn.Message = "O número não é divisivel, por nenhum dos valores"
		jsonreturn.Error = false
		return jsonreturn, nil
	} else {
		jsonreturn.Data = ndiv
		jsonreturn.Message = nil
		jsonreturn.Error = false

		return jsonreturn, nil
	}
}

func ValidaPessoa(jsonreturn models.Jsonreturn, pessoa models.Pessoa) (models.Jsonreturn, error) {

	if pessoa.Name == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Nome não definido"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Nome indefinido")
	} else if pessoa.Age <= 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Idade Ínvalida"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Idade igual a 0 ou menor")
	} else if strings.TrimSpace(strings.ToLower(pessoa.Sexo)) != "masculino" && strings.TrimSpace(strings.ToLower(pessoa.Sexo)) != "feminino" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Sexo Ínvalido! Defina Masculino ou Feminino"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Sexo não existente")
	}

	if pessoa.Age < 25 && strings.TrimSpace(strings.ToLower(pessoa.Sexo)) == "feminino" {
		jsonreturn.Data = (pessoa.Name + "Aceita")
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else {
		jsonreturn.Data = (pessoa.Name + "Não aceita")
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	}
}

func ImprimeNumerosDecrescente() {

}

func MesDoAno() {

}

func Biblioteca() {

}

func ImprimeNumeroProduto() {

}

func ImprimeSol() {

}

func PositivosNegativos() {

}

func ImprimePalavra() {

}

func ParImpar() {

}

func AnaliseNumerica() {

}

func IgualaVetores() {

}

func AnosCrescimento() {

}
