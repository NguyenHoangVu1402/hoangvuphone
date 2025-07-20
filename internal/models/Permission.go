package models

import (
	"time"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          uuid.UUID    `gorm:"type:char(36);primary_key" json:"id"`
	Name        string    `gorm:"size:100;uniqueIndex;not null" json:"name"`  
	Slug        string    `gorm:"size:100;uniqueIndex;not null" json:"slug"`  
	Description string    `gorm:"size:255" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}