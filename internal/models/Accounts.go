package models

import (
	"time"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID          uuid.UUID         `gorm:"type:char(36);primary_key" json:"id"`
	Username    string         `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email       string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"size:255;not null" json:"-"` 
	Avatar      string         `gorm:"size:255" json:"avatar,omitempty"`
	Birthday    *time.Time     `gorm:"type:date" json:"birthday,omitempty"` 
	BonusPoints int            `gorm:"default:0" json:"bonusPoints"`         
	Phone       string         `gorm:"size:20;uniqueIndex" json:"phone,omitempty"`
	Address     string         `gorm:"type:text" json:"address,omitempty"` 
	IsVerified  bool           `gorm:"default:false" json:"isVerified"`   
	LastLogin   *time.Time     `json:"lastLogin,omitempty"`
	Roles       []Role         `gorm:"many2many:account_roles;" json:"roles,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}