package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samratsuzil/api/restapi"
	"github.com/samratsuzil/api/models"
	"github.com/samratsuzil/api/database"
)

func main(){

	fmt.Println("Connecting to DB")
	db := database.ConnectDB()

	fmt.Println("Automigrating User")
	db.AutoMigrate(&models.User{},&models.Book{},&models.Category{})

	fmt.Println("Defer Close DB Connection")
	defer db.Close()

	router := gin.Default()
	r := router.Group("/v1")
	restapi.InitializeRoutes(r)

	fmt.Println("Running API")
	router.Run(":8080")
}