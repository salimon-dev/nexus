package types

import (
	"time"
)

type UserStatus int8

const (
	UserStatusPending  UserStatus = 1
	UserStatusActive   UserStatus = 2
	UserStatusInActive UserStatus = 3
)

func UserStatusToString(status UserStatus) string {
	switch status {
	case UserStatusActive:
		return "active"
	case UserStatusPending:
		return "pending"
	case UserStatusInActive:
		return "inactive"
	default:
		return "none"
	}
}

type UserRole int8

const (
	UserRoleKeyMaker UserRole = 1
	UserRoleAdmin    UserRole = 2
	UserRoleMember   UserRole = 3
)

func UserRoleToString(role UserRole) string {
	switch role {
	case UserRoleKeyMaker:
		return "keymaker"
	case UserRoleAdmin:
		return "admin"
	case UserRoleMember:
		return "member"
	default:
		return "none"
	}
}

type User struct {
	Id           string     `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Credit       int32      `json:"credit"`
	Usage        int32      `json:"usage"`
	Role         UserRole   `json:"role"`
	Status       UserStatus `json:"status"`
	RegisteredAt time.Time  `json:"registered_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type PublicUser struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Credit       int32     `json:"credit"`
	Usage        int32     `json:"usage"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	RegisteredAt time.Time `json:"registered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
