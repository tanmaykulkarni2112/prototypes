package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tanmaykulkarni2112/prototypes/golangdb/model"
)

// intern type shit
func GetHello(c *gin.Context) {
	msg := model.ResponseMsg{
		Message : "Hello",
	}
	c.IndentedJSON(200, msg)
}