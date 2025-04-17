package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/viktoralyoshin/nookly/api-gateway/protos/github.com/viktoralyoshin/nookly/api-gateway/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	app := fiber.New()

	conn, err := grpc.NewClient("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}

	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{UserId: id})
		if err != nil {
			return c.Status(500).SendString("gRPC error: " + err.Error())
		}

		return c.JSON(fiber.Map{
			"id":    res.UserId,
			"name":  res.Name,
			"email": res.Email,
		})
	})

	log.Println("API Gateway running on :8080")
	log.Fatal(app.Listen(":8080"))

}
