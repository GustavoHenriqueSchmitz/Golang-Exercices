package models

type City struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	State_id int    `json:"state_id"`
}
