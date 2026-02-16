package model

import (
	"time"

	"github.com/google/uuid"
)

type Market struct {
	ID           uuid.UUID
	Name         string
	Enabled      bool
	AllowedRoles []string
	CreatedAt    *time.Time
	UpdateAt     *time.Time
	DeletedAt    *time.Time
}
