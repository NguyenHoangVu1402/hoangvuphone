package models

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"uniqueIndex;not_null" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not_null" json:"slug"`
	Description string    `json:"description"`
}