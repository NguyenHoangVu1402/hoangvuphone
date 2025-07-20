package models

import (
	"time"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID         `gorm:"type:char(36);primary_key" json:"id"`
	Name        string         `gorm:"size:50;uniqueIndex;not null" json:"name"`  
	Slug        string         `gorm:"size:50;uniqueIndex;not null" json:"slug"` 
	Description string         `gorm:"size:255" json:"description,omitempty"`
	Level       int            `gorm:"default:1" json:"level"` 
	Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
	Accounts    []Account      `gorm:"many2many:account_roles;" json:"-"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}