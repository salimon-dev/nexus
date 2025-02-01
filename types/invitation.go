package types

import (
	"time"

	"github.com/google/uuid"
)

type InvitationStatus int8

const (
	InvitationStatusActive   InvitationStatus = 1
	InvitationStatusInActive InvitationStatus = 2
)

func InvitationStatusToString(status InvitationStatus) string {
	switch status {
	case InvitationStatusActive:
		return "active"
	case InvitationStatusInActive:
		return "inactive"
	default:
		return "none"
	}
}

type Invitation struct {
	Id             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedBy      uuid.UUID `json:"created_by" gorm:"type:uuid;not null"`
	Code           string    `json:"code" gorm:"type:string;size:12;not null"`
	UsageRemaining int16     `json:"usage_remaining" gorm:"type:numeric;not null"`
	ExpiresAt      time.Time `json:"expires_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
}
