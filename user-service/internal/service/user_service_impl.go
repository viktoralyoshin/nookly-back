package service

import (
	"context"
	"errors"

	"github.com/viktoralyoshin/nookly/user-serivce/internal/model"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/repository"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/utils"
)

var (
	ErrEmailTaken = errors.New("email already taken")
	ErrNameTaken  = errors.New("name already taken")
)

type Service struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &Service{repository: repository}
}

func (s *Service) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, candidate model.CreateUser) (*model.User, error) {
	existsByEmail, _ := s.repository.GetUserByEmail(ctx, candidate.Email)
	if existsByEmail != nil {
		return nil, ErrEmailTaken
	}

	existsByName, _ := s.repository.GetUserByName(ctx, candidate.Name)
	if existsByName != nil {
		return nil, ErrNameTaken
	}

	hashedPassword, err := utils.HashPassword(candidate.Password)
	if err != nil {
		return nil, err
	}

	user := &model.CreateUser{
		Name:     candidate.Name,
		Email:    candidate.Email,
		Password: hashedPassword,
	}

	createUserResponse, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createUserResponse, nil
}
