package types

import (
	"time"

	"github.com/google/uuid"
)

type EntityStatus int8

const (
	EntityStatusActive   EntityStatus = 1
	EntityStatusInActive EntityStatus = 2
)

func EntityStatusToString(status EntityStatus) string {
	switch status {
	case EntityStatusActive:
		return "active"
	case EntityStatusInActive:
		return "inactive"
	default:
		return "none"
	}
}

type Entity struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" gorm:"size:32;unique;not null"`
	Description string    `json:"description" gorm:"size:256"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
}
