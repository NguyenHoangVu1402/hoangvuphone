package repositories

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"hoangvuphone/internal/models"
)

type RoleRepository interface {
	GetAllRoles(search string, page, limit int) ([]models.Role, int64, error)
	GetByID(id uuid.UUID) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetBySlug(slug string) (*models.Role, error)
	Create(role models.Role) (*models.Role, error)
	Update(role *models.Role) (*models.Role, error)
	Delete(id uuid.UUID) error
	AssignPermissions(roleID uuid.UUID, permissionIDs []uuid.UUID) error
	SyncPermissions(roleID uuid.UUID, permissionIDs []uuid.UUID) error
	GetPermissions(roleID uuid.UUID) ([]models.Permission, error)
	CountAccountsByRoleID(roleID uuid.UUID) (int64, error)
	Search(query string) ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (rr *roleRepository) GetAllRoles(search string, page, limit int) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	query := rr.db.Model(&models.Role{})

	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(slug) LIKE ?", searchTerm, searchTerm)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (rr *roleRepository) GetByID(id uuid.UUID) (*models.Role, error) {
	var role models.Role
	err := rr.db.First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *roleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := rr.db.First(&role, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *roleRepository) GetBySlug(slug string) (*models.Role, error) {
	var role models.Role
	err := rr.db.First(&role, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *roleRepository) Create(role models.Role) (*models.Role, error) {
	err := rr.db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *roleRepository) Update(role *models.Role) (*models.Role, error) {
	err := rr.db.Save(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (rr *roleRepository) Delete(id uuid.UUID) error {
	return rr.db.Delete(&models.Role{}, "id = ?", id).Error
}

func (rr *roleRepository) AssignPermissions(roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	role := models.Role{ID: roleID}
	return rr.db.Model(&role).Association("Permissions").Append(permissionIDs)
}

func (rr *roleRepository) SyncPermissions(roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	role := models.Role{ID: roleID}
	return rr.db.Model(&role).Association("Permissions").Replace(permissionIDs)
}

func (rr *roleRepository) GetPermissions(roleID uuid.UUID) ([]models.Permission, error) {
	var role models.Role
	err := rr.db.Preload("Permissions").First(&role, "id = ?", roleID).Error
	if err != nil {
		return nil, err
	}
	return role.Permissions, nil
}

func (rr *roleRepository) CountAccountsByRoleID(roleID uuid.UUID) (int64, error) {
	var count int64
	err := rr.db.Model(&models.AccountRoles{}).Where("role_id = ?", roleID).Count(&count).Error
	return count, err
}

func (rr *roleRepository) Search(query string) ([]models.Role, error) {
	var roles []models.Role
	searchTerm := "%" + strings.ToLower(query) + "%"
	err := rr.db.Where("LOWER(name) LIKE ? OR LOWER(slug) LIKE ?", searchTerm, searchTerm).
		Limit(10).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}