package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDB() {

	connect := "host=localhost user=egadeduc_courses password=egadeduc_courses dbname=egadeduc_courses port=15433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connect))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados", DB)
	}
}
