package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hoangvuphone/internal/models"
)

type PermissionRepository interface {
	GetAll() ([]models.Permission, error)
	GetByID(id uuid.UUID) (*models.Permission, error)
	GetByName(name string) (*models.Permission, error)
	GetBySlug(slug string) (*models.Permission, error)
	Create(permission models.Permission) (*models.Permission, error)
	Update(permission *models.Permission) (*models.Permission, error)
	Delete(id uuid.UUID) error
	GetByRoleID(roleID uuid.UUID) ([]models.Permission, error)
	CountRolesByPermissionID(permissionID uuid.UUID) (int64, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (pr *permissionRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := pr.db.Find(&permissions).Error
	return permissions, err
}

func (pr *permissionRepository) GetByID(id uuid.UUID) (*models.Permission, error) {
	var permission models.Permission
	err := pr.db.First(&permission, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (pr *permissionRepository) GetByName(name string) (*models.Permission, error) {
	var permission models.Permission
	err := pr.db.First(&permission, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (pr *permissionRepository) GetBySlug(slug string) (*models.Permission, error) {
	var permission models.Permission
	err := pr.db.First(&permission, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (pr *permissionRepository) Create(permission models.Permission) (*models.Permission, error) {
	err := pr.db.Create(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (pr *permissionRepository) Update(permission *models.Permission) (*models.Permission, error) {
	err := pr.db.Save(permission).Error
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (pr *permissionRepository) Delete(id uuid.UUID) error {
	return pr.db.Delete(&models.Permission{}, "id = ?", id).Error
}

func (pr *permissionRepository) GetByRoleID(roleID uuid.UUID) ([]models.Permission, error) {
	var role models.Role
	err := pr.db.Preload("Permissions").First(&role, "id = ?", roleID).Error
	if err != nil {
		return nil, err
	}
	return role.Permissions, nil
}

func (pr *permissionRepository) CountRolesByPermissionID(permissionID uuid.UUID) (int64, error) {
	var count int64
	err := pr.db.Model(&models.RolePermissions{}).Where("permission_id = ?", permissionID).Count(&count).Error
	return count, err
}