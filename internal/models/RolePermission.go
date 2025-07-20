package models

import (
	"github.com/google/uuid"
)

type RolePermissions struct {
	RoleID       uuid.UUID `gorm:"type:char(36);primaryKey"`
	PermissionID uuid.UUID `gorm:"type:char(36);primaryKey"`
}