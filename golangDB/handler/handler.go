package handler

import (
	"log"
	"os"

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

// we will need a function with signature function name (c *gin.Context)
// others will not be accepted
func ProcessJson(filePath string) ([]uint8,error) {
	content , err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return content, nil
}

func GetContent(c *gin.Context) {
	resp , err := ProcessJson("content.json")
	if err != nil {
		log.Fatal(err)
	}
	// THIS DOES NOT WORK THE WAY WE WANT
	// c.JSON(200, resp)

	// THIS WORKS
	c.Data(200, "application/json" , resp)

}
