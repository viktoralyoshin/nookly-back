package model

import (
	"time"

	"github.com/viktoralyoshin/nookly/user-serivce/github.com/viktoralyoshin/nookly/user-service/userpb"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

func (r UserRole) ToProto() userpb.UserRoleProto {
	switch r {
	case RoleAdmin:
		return userpb.UserRoleProto_ROLE_ADMIN
	case RoleUser:
		return userpb.UserRoleProto_ROLE_USER
	default:
		return userpb.UserRoleProto_ROLE_USER
	}
}

func FromProto(protoRole userpb.UserRoleProto) UserRole {
	switch protoRole {
	case userpb.UserRoleProto_ROLE_ADMIN:
		return RoleAdmin
	case userpb.UserRoleProto_ROLE_USER:
		return RoleUser
	default:
		return RoleUser
	}
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
}
