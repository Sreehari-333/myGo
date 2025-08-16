package main

import (
	"github.com/gin-gonic/gin"
	"weather-checker/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/weather", controllers.GetWeather)

	r.Run(":8080")
}
