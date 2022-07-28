package models

type Country struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}
