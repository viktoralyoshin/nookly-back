package repository

import (
	"context"

	"github.com/viktoralyoshin/nookly/user-serivce/internal/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.CreateUser) (*model.User, error)
}
