package models

import (
	"time"
	"github.com/google/uuid"
)

type AccountRoles struct {
	AccountID uuid.UUID `gorm:"type:char(36);primaryKey"`
	RoleID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}