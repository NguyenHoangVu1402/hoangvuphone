package dtos

import (
	"github.com/google/uuid"
	"hoangvuphone/internal/models"
	"time"
	
)

type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Slug        string `json:"slug" validate:"required,min=3,max=100,slug"`
	Description string `json:"description,omitempty" validate:"max=255"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Slug        string `json:"slug,omitempty" validate:"omitempty,min=3,max=100,slug"`
	Description string `json:"description,omitempty" validate:"omitempty,max=255"`
}

type PermissionResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToPermissionResponse(permission models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Slug:        permission.Slug,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}