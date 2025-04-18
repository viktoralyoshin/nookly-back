package router

import (
	"github.com/gofiber/fiber/v2"
	userhandler "github.com/viktoralyoshin/nookly/api-gateway/internal/handler/user_handler"
)

func RegisterUserRoutes(group fiber.Router) {
	users := group.Group("/users")
	users.Get("/:id", userhandler.GetUser)
	users.Post("/", userhandler.CreateUser)
}
