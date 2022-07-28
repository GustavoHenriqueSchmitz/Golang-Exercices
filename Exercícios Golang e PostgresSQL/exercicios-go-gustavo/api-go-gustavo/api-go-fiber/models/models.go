package models

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Permissao bool   `json:"permissao"`
}

type Nums struct {
	N1 int `json:"num1"`
	N2 int `json:"num2"`
}

type Jsonreturn struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}

type Pessoa struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sexo string `json:"sexo"`
}

type Nums3 struct {
	N1 int `json:"num1"`
	N2 int `json:"num2"`
	N3 int `json:"num3"`
}

type Biblioteca struct {
	User          string `json:"user"`
	UserType      string `json:"usertype"`
	Book          string `json:"book"`
	DevolucaoDias int    `json:"devolucao"`
}

type NumeProduto struct {
	Num     int `json:"num"`
	Produto int `json:"produto"`
}

type Nums5 struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
	Num3 int `json:"num3"`
	Num4 int `json:"num4"`
	Num5 int `json:"num5"`
}

type PositivoNegativo struct {
	QuantNegativos int `json:"quantidade_negativos"`
	SomaPositivos  int `json:"soma_positivos"`
}

type Palavra struct {
	Palavra string `json:"palavra"`
}

type NumsList struct {
	Numeros []int `json:"list"`
}

type AnaliseNumerica struct {
	Maior           int     `json:"maior"`
	Menor           int     `json:"menor"`
	PercentualPares int     `json:"percentualpares"`
	Media           float64 `json:"media"`
}

type Nums2Lists struct {
	Numeros1 []int `json:"list1"`
	Numeros2 []int `json:"list2"`
}
