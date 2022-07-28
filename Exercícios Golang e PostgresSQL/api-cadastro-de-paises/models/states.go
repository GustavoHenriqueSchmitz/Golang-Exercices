package models

type States struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Acronym    string `json:"acronym"`
	Country_id int    `json:"country_id"`
}
