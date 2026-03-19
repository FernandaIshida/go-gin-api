package database

import (
	"log"

	"fmt"

	"github.com/FernandaIshida/go-gin-api.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	stringConnection := "host=127.0.0.1 user=root password=root dbname=root port=5432 sslmode=disable"
	fmt.Println("DSN:", stringConnection)

	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Student{})
}
