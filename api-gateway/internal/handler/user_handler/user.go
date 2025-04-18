package userhandler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/viktoralyoshin/nookly/api-gateway/internal/grpc"
	"github.com/viktoralyoshin/nookly/api-gateway/internal/logger"
	"github.com/viktoralyoshin/nookly/api-gateway/protos/github.com/viktoralyoshin/nookly/user-service/userpb"
)

// GET /api/users/:id
func GetUser(c *fiber.Ctx) error {
	log := logger.GetLogger()
	id := c.Params("id")

	log.Info().
		Str("method", "GetUser").
		Str("user_id", id).
		Msg("Запрос пользователя")

	res, err := grpc.UserClient.GetUser(context.Background(), &userpb.GetUserRequest{UserId: id})
	if err != nil {
		log.Error().
			Err(err).
			Str("user_id", id).
			Msg("Ошибка при получении пользователя")

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().
		Str("user_id", id).
		Msg("Успешно получен пользователь")

	return c.Status(200).JSON(res)
}

// POST /api/users
func CreateUser(c *fiber.Ctx) error {
	log := logger.GetLogger()
	var body struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		log.Error().
			Err(err).
			Msg("Невалидный payload при создании пользователя")

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Info().
		Str("method", "CreateUser").
		Str("email", body.Email).
		Msg("Создание нового пользователя")

	res, err := grpc.UserClient.CreateUser(context.Background(), &userpb.CreateUserRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("email", body.Email).
			Msg("Ошибка при создании пользователя")

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().
		Str("email", body.Email).
		Msg("Успешно создан пользователь")

	return c.Status(201).JSON(res)
}