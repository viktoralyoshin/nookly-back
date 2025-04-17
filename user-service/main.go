package main

import (
	"context"
	"log"
	"net"

	"github.com/viktoralyoshin/nookly/user-serivce/github.com/viktoralyoshin/nookly/user-service/userpb"
	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		UserId: req.UserId,
		Name: "username",
		Email: "user@gmail.com",
	}, nil
}

func main() {

	PORT := ":50051"

	lis, _ := net.Listen("tcp", PORT)
	
	grpc_s := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpc_s, &server{})
	log.Printf("gRPC UserSevice runnig on %s\n", PORT)

	grpc_s.Serve(lis)
}