package service

import (
	"context"
	"go-grpc-demo/models"
	"go-grpc-demo/pb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Collection *mongo.Collection
}

func NewUserServer(col *mongo.Collection) *UserServer {
	return &UserServer{Collection: col}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	log.Println("Creating user:", req.Name+" "+req.Email)

	user := models.User{Name: req.Name, Email: req.Email}
	res, err := s.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return &pb.User{Id: oid.Hex(), Name: user.Name, Email: user.Email}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {

	oid, _ := primitive.ObjectIDFromHex(req.Id)
	var user models.User
	err := s.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &pb.User{Id: user.ID.Hex(), Name: user.Name, Email: user.Email}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	cursor, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for cursor.Next(ctx) {
		var u models.User
		cursor.Decode(&u)
		users = append(users, &pb.User{Id: u.ID.Hex(), Name: u.Name, Email: u.Email})
	}
	return &pb.ListUsersResponse{Users: users}, nil
}
