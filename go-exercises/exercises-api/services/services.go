package services

import (
	"errors"
	"exercises-api/models"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func ValidaIdade(jsonreturn models.Jsonreturn, user models.User) (models.Jsonreturn, error) {

	user.Name = strings.TrimSpace(strings.Title(user.Name))

	if user.Name == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Nome Indefinido"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Nome indefinido")
	} else if user.Age < 0 {
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

	pessoa.Name = strings.TrimSpace(strings.Title(pessoa.Name))
	pessoa.Sexo = strings.TrimSpace(strings.ToLower(pessoa.Sexo))

	if pessoa.Name == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Nome não definido"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Nome indefinido")
	} else if pessoa.Age < 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Idade Ínvalida"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Idade menor que 0")
	} else if pessoa.Sexo != "masculino" && pessoa.Sexo != "feminino" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Sexo Ínvalido! Defina Masculino ou Feminino"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Sexo não existente")
	}

	if pessoa.Age < 25 && pessoa.Sexo == "feminino" {
		jsonreturn.Data = (pessoa.Name + " Aceita")
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else {
		jsonreturn.Data = (pessoa.Name + " Não aceita")
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	}
}

func ImprimeNumerosDecrescente(jsonreturn models.Jsonreturn, nums models.Nums3) models.Jsonreturn {
	numslist := []int{}

	numslist = append(numslist, nums.N1)
	numslist = append(numslist, nums.N2)
	numslist = append(numslist, nums.N3)

	sort.Sort(sort.Reverse(sort.IntSlice(numslist)))

	jsonreturn.Data = numslist
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn

}

func MesDoAno(jsonreturn models.Jsonreturn, num models.Nums) (models.Jsonreturn, error) {

	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

	if num.N1 < 1 || num.N1 > 12 {
		jsonreturn.Data = nil
		jsonreturn.Message = "O mês digitado não existe"
		jsonreturn.Error = false
		return jsonreturn, errors.New("Erro! Mês Inesistente")
	}

	jsonreturn.Data = meses[num.N1-1]
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func Biblioteca(jsonreturn models.Jsonreturn, emprestimo models.Biblioteca) (models.Jsonreturn, error) {

	emprestimo.User = strings.TrimSpace(strings.Title(emprestimo.User))
	emprestimo.Book = strings.TrimSpace(strings.Title(emprestimo.Book))
	emprestimo.UserType = strings.TrimSpace(strings.ToLower(emprestimo.UserType))

	if emprestimo.User == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Usuário não informado"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Usuário não informado")
	} else if emprestimo.Book == "" {
		jsonreturn.Data = nil
		jsonreturn.Message = "Livro não informado"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Livro não informado")
	}

	if emprestimo.UserType == "professor" {
		emprestimo.DevolucaoDias = 10
		jsonreturn.Data = emprestimo
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else if emprestimo.UserType == "aluno" {
		emprestimo.DevolucaoDias = 3
		jsonreturn.Data = emprestimo
		jsonreturn.Message = nil
		jsonreturn.Error = false
		return jsonreturn, nil
	} else {
		jsonreturn.Data = nil
		jsonreturn.Message = "Tipo de usuário ínvalido! Digite professor ou aluno"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Tipo de usuário ínvalido")
	}
}

func ImprimeNumeroProduto(jsonreturn models.Jsonreturn, num models.Nums) (models.Jsonreturn, error) {

	result_list := []models.NumeProduto{}

	if num.N1 < 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Números abaixo de 0, não são validos"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Número digitado abaixo de 0")
	}

	for cont := 0; cont <= num.N1; cont += 1 {
		num_Produto := models.NumeProduto{Num: cont, Produto: num.N1 * cont}
		result_list = append(result_list, num_Produto)
	}

	jsonreturn.Data = result_list
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func ImprimeSol(jsonreturn models.Jsonreturn, num int) (models.Jsonreturn, error) {

	if num < 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Números abaixo de 0, não são aceitos"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Número abaixo de 0")
	}

	sol_list := []string{}

	for cont := 0; cont < num; cont += 1 {
		sol_list = append(sol_list, "SOL")
	}

	jsonreturn.Data = sol_list
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func PositivosNegativos(jsonreturn models.Jsonreturn, numeros models.Nums5) models.Jsonreturn {

	nums_negativos := 0
	soma_positivos := 0

	v := reflect.ValueOf(numeros)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	for _, value := range values {

		if value.(int) < 0 {
			nums_negativos += 1
		} else {
			soma_positivos += value.(int)
		}
	}

	resul := models.PositivoNegativo{}
	resul.QuantNegativos = nums_negativos
	resul.SomaPositivos = soma_positivos

	jsonreturn.Data = resul
	jsonreturn.Message = nil
	jsonreturn.Error = false

	return jsonreturn
}

func Tabuada(jsonreturn models.Jsonreturn, num models.Nums) (models.Jsonreturn, error) {

	if num.N1 > 10 || num.N1 < 0 {
		jsonreturn.Data = nil
		jsonreturn.Message = "Número Ínvalido! Digite um valor de 0 a 10"
		jsonreturn.Error = true
		return jsonreturn, errors.New("Erro! Valor maior que 10, ou menor que 0")
	}

	tabuada_result := []int{}
	for cont := 0; cont <= 10; cont += 1 {
		tabuada_result = append(tabuada_result, num.N1*cont)
	}

	jsonreturn.Data = tabuada_result
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn, nil
}

func ImprimePalavra(jsonreturn models.Jsonreturn, palavra models.Palavra) models.Jsonreturn {

	palavra.Palavra = strings.ToUpper(strings.TrimSpace(palavra.Palavra))

	list_result := []string{}

	for cont_letra := 1; cont_letra <= len(palavra.Palavra); cont_letra += 1 {
		result := ""
		for cont_palavra := 0; cont_palavra < cont_letra; cont_palavra += 1 {
			result += " " + palavra.Palavra + " "
		}
		list_result = append(list_result, result)

	}

	jsonreturn.Data = list_result
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn
}

func ParImpar(jsonreturn models.Jsonreturn, numeros models.NumsList) models.Jsonreturn {

	par_impar_list := []string{}

	for _, num := range numeros.Numeros {
		result := ""
		if num%2 == 0 {
			result = strconv.Itoa(num) + " Par"
		} else {
			result = strconv.Itoa(num) + " Ímpar"
		}

		par_impar_list = append(par_impar_list, result)
	}

	jsonreturn.Data = par_impar_list
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn
}

func AnaliseNumerica(jsonreturn models.Jsonreturn, numeros models.NumsList) models.Jsonreturn {

	resultado_analise := models.AnaliseNumerica{}
	quantidade_npares := 0
	soma := 0

	for _, num := range numeros.Numeros {

		if num > resultado_analise.Maior {
			resultado_analise.Maior = num
		} else if num < resultado_analise.Menor || resultado_analise.Menor == 0 {
			resultado_analise.Menor = num
		}

		if num%2 == 0 {
			quantidade_npares += 1
		}

		soma += num
	}

	resultado_analise.PercentualPares = (quantidade_npares * 10) / 1
	resultado_analise.Media = float64(soma) / 10

	jsonreturn.Data = resultado_analise
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn
}

func IgualaVetores(jsonreturn models.Jsonreturn, vetores models.Nums2Lists) models.Jsonreturn {

	result := []int{}
	quebra_while := 0

	for _, nv1 := range vetores.Numeros1 {
		for _, nv2 := range vetores.Numeros2 {
			if nv1 == nv2 {
				for _, nvresul := range result {
					if nv1 == nvresul {
						quebra_while = 1
						break
					}
				}

				if quebra_while > 0 {
					break
				}

				result = append(result, nv1)
				break
			}
		}
	}

	jsonreturn.Data = result
	jsonreturn.Message = nil
	jsonreturn.Error = false

	return jsonreturn
}

func AnosCrescimento(jsonreturn models.Jsonreturn) models.Jsonreturn {

	chico := 1.50
	juca := 1.10
	anos := 0

	for juca < chico {

		chico += 0.02
		juca += 0.03
		anos += 1
	}

	jsonreturn.Data = anos
	jsonreturn.Message = nil
	jsonreturn.Error = false
	return jsonreturn
}
