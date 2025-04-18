package grpc

import (
	"log"
	"os"

	"github.com/viktoralyoshin/nookly/api-gateway/protos/github.com/viktoralyoshin/nookly/user-service/userpb"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient userpb.UserServiceClient
)

func InitGRPC() {
	userServiceAddress := os.Getenv("USER_SERVICE_ADDRESS")

	userServiceConn, err := grpc.NewClient(userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}

	UserClient = userpb.NewUserServiceClient(userServiceConn)
}
