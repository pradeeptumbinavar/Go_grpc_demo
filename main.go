package main

import (
	"go-grpc-demo/db"
	"go-grpc-demo/pb"
	"go-grpc-demo/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	collection := db.ConnectMongo("mongodb://localhost:27017")

	grpcServer := grpc.NewServer()
	userService := service.NewUserServer(collection)
	pb.RegisterUserServiceServer(grpcServer, userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server running on port 50051")
	grpcServer.Serve(lis)
}
