package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello",
	})
}

func main() {

	router := gin.Default()

	router.GET("/", helloHandler)

	router.Run()
}
