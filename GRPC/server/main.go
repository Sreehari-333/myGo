package main

import (
	"context"
	"fmt"
	authpb "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type authServer struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *authServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {

	fmt.Println("Signup called : ", req.Email)

	return &authpb.SignupResponse{
		UserId:  "12345",
		Message: "Signup successfull",
	}, nil
}

func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	fmt.Println("Signup called : ", req.Email)

	return &authpb.LoginResponse{
		Token:   "token",
		Message: "Login successfull",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(grpcServer, &authServer{})

	log.Println("Auth gRPC Server started on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
