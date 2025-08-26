package main

import (
	"context"
	authpb "grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connnect %v", err)
	}

	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.Signup(ctx, &authpb.SignupRequest{
		Email:    "test@gmail.com",
		Password: "12345",
	})

	if err != nil {
		log.Fatalf("Signup failed: %v", err)
	}

	log.Println("Signup response :", res)

	loginRes, err := client.Login(ctx, &authpb.LoginRequest{
		Email:    "test@gmail.com",
		Password: "12345",
	})

	if err != nil {
		log.Fatalf("failed to login : %v", err)
	}

	log.Println("login response : ", loginRes)
}
