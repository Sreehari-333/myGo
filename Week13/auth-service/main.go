package main

import (
	"auth/db"
	"auth/handlers"
	"auth/middleware"
	"auth/models"
	pb "auth/proto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

// gRPC method
func (s *authServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Normally fetch from DB
	return &pb.GetUserResponse{
		UserId: req.UserId,
		Email:  "test@example.com",
	}, nil
}

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	//  Start gRPC server in goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterAuthServiceServer(grpcServer, &authServer{})

		log.Println("Auth gRPC server running on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//  Start HTTP server
	r := gin.Default()
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// Protected route
	protected := r.Group("/user")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetString("user_id")
			c.JSON(200, gin.H{
				"message": "Welcome to your profile",
				"user_id": userID,
			})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default if not set
	}
	log.Println("Auth REST server running on :" + port)
	r.Run(":" + port)
}
