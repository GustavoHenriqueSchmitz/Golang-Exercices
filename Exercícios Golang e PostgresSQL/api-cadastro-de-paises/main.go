package main

import (
	"api-cadastro-de-paises/database"
	"api-cadastro-de-paises/routes"
)

func main() {
	database.ConectDB()
	routes.Requests()
}
