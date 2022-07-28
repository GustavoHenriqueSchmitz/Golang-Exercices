package models

type Jsonreturn struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}
