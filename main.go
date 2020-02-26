package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samratsuzil/api/restapi"
)

func main(){
	fmt.Println("Hello")

	router := gin.Default()
	r := router.Group("/v1")
	restapi.InitializeRoutes(r)
	router.Run()
}