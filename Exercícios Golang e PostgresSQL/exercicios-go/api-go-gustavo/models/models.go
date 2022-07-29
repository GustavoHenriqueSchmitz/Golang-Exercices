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
