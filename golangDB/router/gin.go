package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tanmaykulkarni2112/prototypes/golangdb/handler"
)

func CreateGinRouter() {
	router := gin.Default()
	router.GET("/hello", handler.GetHello)
	err := router.Run("localhost:9991")
	if err != nil {
		log.Fatal("Router could not be started")
	}
}