package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"hoangvuphone/internal/dtos"
	"hoangvuphone/internal/models"
	"hoangvuphone/internal/repositories"
)

type PermissionService interface {
	GetAllPermissions() ([]dtos.PermissionResponse, error)
	GetPermissionByID(id uuid.UUID) (*dtos.PermissionResponse, error)
	CreatePermission(input dtos.CreatePermissionRequest) (*dtos.PermissionResponse, error)
	UpdatePermission(id uuid.UUID, input dtos.UpdatePermissionRequest) (*dtos.PermissionResponse, error)
	DeletePermission(id uuid.UUID) error
	GetPermissionsByRole(roleID uuid.UUID) ([]dtos.PermissionResponse, error)
}

type permissionService struct {
	permissionRepo repositories.PermissionRepository
}

func NewPermissionService(permissionRepo repositories.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func (ps *permissionService) GetAllPermissions() ([]dtos.PermissionResponse, error) {
	permissions, err := ps.permissionRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []dtos.PermissionResponse
	for _, perm := range permissions {
		response = append(response, dtos.ToPermissionResponse(perm))
	}

	return response, nil
}

func (ps *permissionService) GetPermissionByID(id uuid.UUID) (*dtos.PermissionResponse, error) {
	permission, err := ps.permissionRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("permission not found")
		}
		return nil, err
	}

	response := dtos.ToPermissionResponse(*permission)
	return &response, nil
}

func (ps *permissionService) CreatePermission(input dtos.CreatePermissionRequest) (*dtos.PermissionResponse, error) {
	// Check if permission with same name or slug already exists
	if _, err := ps.permissionRepo.GetByName(input.Name); err == nil {
		return nil, errors.New("permission with this name already exists")
	}

	if _, err := ps.permissionRepo.GetBySlug(input.Slug); err == nil {
		return nil, errors.New("permission with this slug already exists")
	}

	permission := models.Permission{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
	}

	createdPerm, err := ps.permissionRepo.Create(permission)
	if err != nil {
		return nil, err
	}

	response := dtos.ToPermissionResponse(*createdPerm)
	return &response, nil
}

func (ps *permissionService) UpdatePermission(id uuid.UUID, input dtos.UpdatePermissionRequest) (*dtos.PermissionResponse, error) {
	permission, err := ps.permissionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if name is being changed and already exists
	if input.Name != "" && input.Name != permission.Name {
		if _, err := ps.permissionRepo.GetByName(input.Name); err == nil {
			return nil, errors.New("permission with this name already exists")
		}
		permission.Name = input.Name
	}

	// Check if slug is being changed and already exists
	if input.Slug != "" && input.Slug != permission.Slug {
		if _, err := ps.permissionRepo.GetBySlug(input.Slug); err == nil {
			return nil, errors.New("permission with this slug already exists")
		}
		permission.Slug = input.Slug
	}

	if input.Description != "" {
		permission.Description = input.Description
	}

	updatedPerm, err := ps.permissionRepo.Update(permission)
	if err != nil {
		return nil, err
	}

	response := dtos.ToPermissionResponse(*updatedPerm)
	return &response, nil
}

func (ps *permissionService) DeletePermission(id uuid.UUID) error {
	// Check if permission is assigned to any role
	count, err := ps.permissionRepo.CountRolesByPermissionID(id)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("cannot delete permission assigned to roles")
	}

	return ps.permissionRepo.Delete(id)
}

func (ps *permissionService) GetPermissionsByRole(roleID uuid.UUID) ([]dtos.PermissionResponse, error) {
	permissions, err := ps.permissionRepo.GetByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	var response []dtos.PermissionResponse
	for _, perm := range permissions {
		response = append(response, dtos.ToPermissionResponse(perm))
	}

	return response, nil
}