package main

import (
	"log"
	"os"
	"product/db"
	"product/handlers"
	"product/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	db.Connect()
	db.DB.AutoMigrate(&models.Product{})

	// Router
	r := gin.Default()
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
