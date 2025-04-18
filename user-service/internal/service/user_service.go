package service

import (
	"context"

	"github.com/viktoralyoshin/nookly/user-serivce/internal/model"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, candidate model.CreateUser) (*model.User, error)
}
