package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"order/db"
	"order/handlers"
	"order/models"
	pb "order/proto"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetUserFromAuth(userID uint) (*pb.GetUserResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect to Auth service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	res, err := client.GetUser(context.Background(), &pb.GetUserRequest{
		UserId: uint32(userID),
	})
	if err != nil {
		log.Printf("error calling GetUser: %v", err)
		return nil, err
	}

	return res, nil
}

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	db.Connect()
	db.DB.AutoMigrate(&models.Order{})

	// Router
	r := gin.Default()
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)
	r.GET("/orders/:id", handlers.GetOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
