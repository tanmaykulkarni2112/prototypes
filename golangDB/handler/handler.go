package handler

import (
	"log"
	"net/http"
	"os"
	"strings"

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

// Create the function to receive the data from the req.body 
// populate inside the map of [string]string
// verify for the tableName 

// we will have to check if the table exists if not we 
// will have to create a file with the tablename

// ----- Needs might need to rename the function to have proper scope

func PostPayload(c *gin.Context) {
	var req model.RequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : err.Error(),
		})
		return
	}

	// check if the file exists.. if it does not exist then create a new file
	fileName := strings.ToLower(req.TableName)
	_ , err := os.Stat(fileName)
	// check if the error is file missing error 
	if os.IsNotExist(err) {
		// The req.data is already JSON.RawMessage in
		// model, that is why we can just pass the data directly 
		writeErr := os.WriteFile(fileName, req.Data, 0666)
		if writeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": writeErr.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "file created",
		})
    return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	} else {
		// file exists
		// would we will apppend the data on the existing json file 
	}	
}