package controllers

import (
	"api-go-fiber/models"
	"api-go-fiber/services"

	"github.com/gofiber/fiber/v2"
)

func Home(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	user := models.User{}
	err := r.BodyParser(&user)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.ValidaIdade(jsonreturn, user)
	return r.JSON(jsonreturn)
}

func Ex01(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	nums := models.Nums{}
	err := r.BodyParser(&nums)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.SomaNumeros(jsonreturn, nums)
	return r.JSON(jsonreturn)
}

func Ex02(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}
	err := r.BodyParser(&num)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.VerificaDivisao(jsonreturn, num)
	return r.JSON(jsonreturn)
}

func Ex03(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	pessoa := models.Pessoa{}
	err := r.BodyParser(&pessoa)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.ValidaPessoa(jsonreturn, pessoa)
	return r.JSON(jsonreturn)
}

func Ex04(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	nums := models.Nums3{}

	err := r.BodyParser(&nums)
	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.ImprimeNumerosDecrescente(jsonreturn, nums)
	return r.JSON(jsonreturn)
}

func Ex05(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}

	err := r.BodyParser(&num)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.MesDoAno(jsonreturn, num)
	return r.JSON(jsonreturn)
}

func Ex06(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	emprestimo := models.Biblioteca{}

	err := r.BodyParser(&emprestimo)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.Biblioteca(jsonreturn, emprestimo)
	return r.JSON(jsonreturn)
}

func Ex07(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}

	err := r.BodyParser(&num)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.ImprimeNumeroProduto(jsonreturn, num)
	return r.JSON(jsonreturn)
}

func Ex08(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}

	err := r.BodyParser(&num)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.ImprimeSol(jsonreturn, num.N1)
	return r.JSON(jsonreturn)
}

func Ex09(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	numeros := models.Nums5{}

	err := r.BodyParser(&numeros)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.PositivosNegativos(jsonreturn, numeros)
	return r.JSON(jsonreturn)
}

func Ex10(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	num := models.Nums{}

	err := r.BodyParser(&num)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn, _ = services.Tabuada(jsonreturn, num)
	return r.JSON(jsonreturn)
}

func Ex11(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	palavra := models.Palavra{}

	err := r.BodyParser(&palavra)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.ImprimePalavra(jsonreturn, palavra)
	return r.JSON(jsonreturn)
}

func Ex12(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	numeros := models.NumsList{}

	err := r.BodyParser(&numeros)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	} else if len(numeros.Numeros) != 15 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Quantidade ínvalida, digite 15 números"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.ParImpar(jsonreturn, numeros)
	return r.JSON(jsonreturn)
}

func Ex13(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	numeros := models.NumsList{}

	err := r.BodyParser(&numeros)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	} else if len(numeros.Numeros) != 10 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Quantidade ínvalida, digite 10 números"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.AnaliseNumerica(jsonreturn, numeros)
	return r.JSON(jsonreturn)
}

func Ex14(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}
	vetores := models.Nums2Lists{}

	err := r.BodyParser(&vetores)

	if err != nil {
		jsonreturn.Data = nil
		jsonreturn.Message = "Ops! Erro ao enviar dados"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	} else if len(vetores.Numeros1) != 5 || len(vetores.Numeros2) != 5 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Quantidade ínvalida, digite 5 números para cada vetor"
		jsonreturn.Error = true
		return r.JSON(jsonreturn)
	}

	jsonreturn = services.IgualaVetores(jsonreturn, vetores)
	return r.JSON(jsonreturn)
}

func Ex15(r *fiber.Ctx) error {
	jsonreturn := models.Jsonreturn{}

	jsonreturn = services.AnosCrescimento(jsonreturn)
	return r.JSON(jsonreturn)
}
