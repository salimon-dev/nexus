package types

import (
	"time"

	"github.com/google/uuid"
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
	Id           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Username     string     `json:"username" gorm:"size:32;unique;not null"`
	Password     string     `json:"password" gorm:"size:32"`
	InvitationId uuid.UUID  `json:"invitation_id" gorm:"type:uuid"`
	Credit       int32      `json:"credit" gorm:"type:numeric"`
	Usage        int32      `json:"usage" gorm:"type:numeric"`
	Role         UserRole   `json:"role" gorm:"type:numeric"`
	Status       UserStatus `json:"status" gorm:"type:numeric"`
	RegisteredAt time.Time  `json:"registered_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
}

type PublicUser struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Credit       int32     `json:"credit"`
	Usage        int32     `json:"usage"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	RegisteredAt time.Time `json:"registered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
