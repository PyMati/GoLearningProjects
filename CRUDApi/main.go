package main

import (
	"github.com/gin-gonic/gin"
)

var db = connectToDb()

func main() {

	defer closeConnection(db)

	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/create-table", createTable)

	r.Run()
}
