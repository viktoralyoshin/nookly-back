package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/viktoralyoshin/nookly/api-gateway/internal/grpc"
	"github.com/viktoralyoshin/nookly/api-gateway/internal/logger"
	"github.com/viktoralyoshin/nookly/api-gateway/internal/router"
)


func main() {

	log := logger.InitLogger("api-gateway")

	grpc.InitGRPC()

	app := fiber.New()

	router.SetupRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start API Gateway")
	}

	log.Info().Msg("API Gateway started on 8080")

}
