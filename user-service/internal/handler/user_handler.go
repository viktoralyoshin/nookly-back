package handler

import (
	"context"

	"github.com/viktoralyoshin/nookly/user-serivce/github.com/viktoralyoshin/nookly/user-service/userpb"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/model"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.service.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		UserId:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role.ToProto(),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.GetUserResponse, error) {

	candidate := &model.CreateUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.service.CreateUser(ctx, *candidate)
	if err != nil {
		return nil, err
	}

	response := &userpb.GetUserResponse{
		UserId: user.ID,
		Email: user.Email,
		Name: user.Name,
		Role: user.Role.ToProto(),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return response, nil

}
