package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "tasrifin"
	DB_PASSWORD = "tasrifin"
	DB_NAME     = "dbtest"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// log.Default().Println("Connection Success")

	// db.Debug().AutoMigrate(models.User{}, models.Product{})
	return db
}
