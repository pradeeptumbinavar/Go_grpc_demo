package main

import (
	"context"
	"log"
	"time"

	"go-grpc-demo/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// CreateUser
	res, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "TestUser3",
		Email: "test3@example.com",
	})
	if err != nil {
		log.Fatal("CreateUser error:", err)
	}

	log.Println("Created:", res)

	// ListUsers
	list, _ := client.ListUsers(ctx, &pb.ListUsersRequest{})
	for _, user := range list.Users {
		log.Println("User:", user)
	}
}
