package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ok",
	})
}

func createTable(c *gin.Context) {
	var table Table

	if err := c.BindJSON(&table); err != nil {
		fmt.Println("An error occured binding table name.")
		c.JSON(400, gin.H{
			"message": "An error occured while binding table name.",
		})
		return
	}

	query := fmt.Sprintf("CREATE TABLE %s(%s)", table.Name, defaultId)
	_, err := db.Exec(query)

	if err != nil {
		fmt.Println("An error occured while creating a table.", err)
		c.JSON(400, gin.H{
			"message": "An error occured while creating table in database.",
		})
		return
	}

	fmt.Println("Table created.", table.Name)
	c.IndentedJSON(http.StatusCreated, table)
}
