package main

import (
	"fmt"
	"log"
	"sesi8-gin/controllers"
	"sesi8-gin/models"
	"sesi8-gin/repositories"
	"sesi8-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "tasrifin"
	DB_PASSWORD = "tasrifin"
	DB_NAME     = "dbtest"
	APP_PORT    = ":8888"
)

func main() {
	db := connectDB()

	router := gin.Default()

	personRepo := repositories.NewPersonRepo(db)
	personService := services.NewPersonServices(personRepo)
	personController := controllers.NewPersonController(personService)

	router.POST("/persons", personController.CreateNewPerson)
	router.GET("/persons", personController.GetAllPersons)

	departmentRepo := repositories.NewDepartmentRepo(db)
	departmentService := services.NewDepartmentServices(departmentRepo)
	departmentController := controllers.NewDepartmentController(departmentService)

	router.POST("/departments", departmentController.CreateNewDepartment)
	router.GET("/departments", departmentController.GetAllDepartments)

	log.Println("Server RUnning at Port : ", APP_PORT)
	router.Run(APP_PORT)

}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	log.Default().Println("Connection to Database is Successfull")
	// db.AutoMigrate(models.Person{})
	db.AutoMigrate(models.Department{})
	return db

}
