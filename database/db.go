package database

import (
	"log"

	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectKeyDB := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectKeyDB))
	if err != nil {
		log.Panic("Erro ao conectar com Banco de Dados!")
	}
	DB.AutoMigrate(&models.User{})
}
