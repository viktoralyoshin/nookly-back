package main

import (
	"log"
	"net"
	"os"

	"github.com/viktoralyoshin/nookly/user-serivce/github.com/viktoralyoshin/nookly/user-service/userpb"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/db"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/handler"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/repository"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/service"
	"google.golang.org/grpc"
)

func main() {


	PORT := os.Getenv("USER_SERVICE_PORT")

	db, err := db.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpc_s := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpc_s, handler)

	log.Printf("gRPC UserSevice runnig on %s\n", PORT)
	if err := grpc_s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
