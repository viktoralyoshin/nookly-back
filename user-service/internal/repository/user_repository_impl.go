package repository

import (
	"context"
	"database/sql"
	"github.com/viktoralyoshin/nookly/user-serivce/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	query := "SELECT id, email, username, role, created_at, updated_at FROM users WHERE id=$1"

	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT id, email, username, role, created_at, updated_at FROM users WHERE email=$1"

	var user model.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	query := "SELECT id, email, username, role, created_at, updated_at FROM users WHERE username=$1"

	var user model.User
	err := r.db.QueryRowContext(ctx, query, name).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *model.CreateUser) (*model.User, error) {
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id, email, username, role, created_at, updated_at"

	var userResponse model.User
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, user.Password).Scan(&userResponse.ID, &userResponse.Email, &userResponse.Name, &userResponse.Role, &userResponse.CreatedAt, &userResponse.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
