package dtos

import (
	"github.com/google/uuid"
	"hoangvuphone/internal/models"
	"time"
)

type CreateRoleRequest struct {
	Name         string      `json:"name" validate:"required,min=3,max=50"`
	Slug         string      `json:"slug" validate:"required,min=3,max=50,slug"`
	Description  string      `json:"description,omitempty" validate:"max=255"`
	Level        int         `json:"level" validate:"required,min=1"`
	PermissionIDs []uuid.UUID `json:"permissionIds,omitempty"`
}

type UpdateRoleRequest struct {
	Name         *string      `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	Slug         *string      `json:"slug,omitempty" validate:"omitempty,min=3,max=50,slug"`
	Description  *string      `json:"description,omitempty" validate:"omitempty,max=255"`
	Level        *int         `json:"level,omitempty" validate:"omitempty,min=1"`
	PermissionIDs []uuid.UUID `json:"permissionIds,omitempty"`
}

type RoleResponse struct {
	ID           uuid.UUID          `json:"id"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Description  string             `json:"description,omitempty"`
	Level        int                `json:"level"`
	Permissions  []PermissionResponse `json:"permissions,omitempty"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}


type PaginatedRoleResponse struct {
	Total int64           `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
	Data  []RoleResponse  `json:"data"`
}

func ToRoleResponse(role models.Role, permissions []models.Permission) RoleResponse {
	var permResponses []PermissionResponse
	for _, perm := range permissions {
		permResponses = append(permResponses, PermissionResponse{
			ID:          perm.ID,
			Name:        perm.Name,
			Slug:        perm.Slug,
			Description: perm.Description,
		})
	}

	return RoleResponse{
		ID:           role.ID,
		Name:         role.Name,
		Slug:         role.Slug,
		Description:  role.Description,
		Level:        role.Level,
		Permissions:  permResponses,
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
	}
}