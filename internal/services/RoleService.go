package services

import (
	"errors"
	

	"hoangvuphone/internal/dtos"
	"hoangvuphone/internal/repositories"
	"hoangvuphone/internal/models"

	"github.com/google/uuid"
	
	"gorm.io/gorm"
)

type RoleService interface {
	GetAllRoles(search string, page, limit int) ([]dtos.RoleResponse, int64, error)
	GetRoleByID(id uuid.UUID) (*dtos.RoleResponse, error)
	CreateRole(input dtos.CreateRoleRequest) (*dtos.RoleResponse, error)
	UpdateRole(id uuid.UUID, input dtos.UpdateRoleRequest) (*dtos.RoleResponse, error)
	DeleteRole(id uuid.UUID) error
	SearchRoles(query string) ([]dtos.RoleResponse, error)
	GetPermissions(roleID uuid.UUID) ([]models.Permission, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
	permissionRepo repositories.PermissionRepository
}

func NewRoleService(roleRepo repositories.RoleRepository, permissionRepo repositories.PermissionRepository) RoleService {
	return &roleService{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}

func (rs *roleService) GetAllRoles(search string, page, limit int) ([]dtos.RoleResponse, int64, error) {
	roles, total, err := rs.roleRepo.GetAllRoles(search, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var response []dtos.RoleResponse
	for _, role := range roles {
		permissions, err := rs.permissionRepo.GetByRoleID(role.ID)
		if err != nil {
			return nil, 0, err
		}

		response = append(response, dtos.ToRoleResponse(role, permissions))
	}

	return response, total, nil
}

func (rs *roleService) GetPermissions(roleID uuid.UUID) ([]models.Permission, error) {
	return rs.permissionRepo.GetByRoleID(roleID)
}


func (rs *roleService) GetRoleByID(id uuid.UUID) (*dtos.RoleResponse, error) {
	role, err := rs.roleRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	permissions, err := rs.permissionRepo.GetByRoleID(role.ID)
	if err != nil {
		return nil, err
	}

	response := dtos.ToRoleResponse(*role, permissions)
	return &response, nil
}

func (rs *roleService) CreateRole(input dtos.CreateRoleRequest) (*dtos.RoleResponse, error) {
	// Check if role with same name or slug already exists
	if _, err := rs.roleRepo.GetByName(input.Name); err == nil {
		return nil, errors.New("role with this name already exists")
	}

	if _, err := rs.roleRepo.GetBySlug(input.Slug); err == nil {
		return nil, errors.New("role with this slug already exists")
	}

	// Create role
	role := models.Role{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		Level:       input.Level,
	}

	createdRole, err := rs.roleRepo.Create(role)
	if err != nil {
		return nil, err
	}

	// Assign permissions
	if len(input.PermissionIDs) > 0 {
		if err := rs.roleRepo.AssignPermissions(createdRole.ID, input.PermissionIDs); err != nil {
			return nil, err
		}
	}

	// Get permissions for response
	permissions, err := rs.permissionRepo.GetByRoleID(createdRole.ID)
	if err != nil {
		return nil, err
	}

	response := dtos.ToRoleResponse(*createdRole, permissions)
	return &response, nil
}

func (rs *roleService) UpdateRole(id uuid.UUID, input dtos.UpdateRoleRequest) (*dtos.RoleResponse, error) {
	// Lấy role hiện tại
	role, err := rs.roleRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Kiểm tra và cập nhật Name nếu thay đổi
	if input.Name != nil && *input.Name != role.Name {
		if _, err := rs.roleRepo.GetByName(*input.Name); err == nil {
			return nil, errors.New("role with this name already exists")
		}
		role.Name = *input.Name
	}

	// Kiểm tra và cập nhật Slug nếu thay đổi
	if input.Slug != nil && *input.Slug != role.Slug {
		if _, err := rs.roleRepo.GetBySlug(*input.Slug); err == nil {
			return nil, errors.New("role with this slug already exists")
		}
		role.Slug = *input.Slug
	}

	// Cập nhật Description nếu có
	if input.Description != nil {
		role.Description = *input.Description
	}

	// Cập nhật Level nếu có
	if input.Level != nil {
		role.Level = *input.Level
	}

	// Cập nhật role
	updatedRole, err := rs.roleRepo.Update(role)
	if err != nil {
		return nil, err
	}

	// Cập nhật lại permissions nếu có gửi lên
	if input.PermissionIDs != nil {
		if err := rs.roleRepo.SyncPermissions(updatedRole.ID, input.PermissionIDs); err != nil {
			return nil, err
		}
	}

	// Lấy lại danh sách permission mới cho response
	permissions, err := rs.permissionRepo.GetByRoleID(updatedRole.ID)
	if err != nil {
		return nil, err
	}

	response := dtos.ToRoleResponse(*updatedRole, permissions)
	return &response, nil
}


func (rs *roleService) DeleteRole(id uuid.UUID) error {
	// Check if role is assigned to any account
	count, err := rs.roleRepo.CountAccountsByRoleID(id)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("cannot delete role assigned to accounts")
	}

	return rs.roleRepo.Delete(id)
}

func (rs *roleService) SearchRoles(query string) ([]dtos.RoleResponse, error) {
	roles, err := rs.roleRepo.Search(query)
	if err != nil {
		return nil, err
	}

	var response []dtos.RoleResponse
	for _, role := range roles {
		permissions, err := rs.permissionRepo.GetByRoleID(role.ID)
		if err != nil {
			return nil, err
		}

		response = append(response, dtos.ToRoleResponse(role, permissions))
	}

	return response, nil
}