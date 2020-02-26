package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/astranix/ims-backend-go/app/restapi"
)

func main(){
	fmt.Println("Hello")

	router := gin.Default()
	r := router.Group("/v1")
	restapi.InitializeRoutes(r)
	router.Run()
}