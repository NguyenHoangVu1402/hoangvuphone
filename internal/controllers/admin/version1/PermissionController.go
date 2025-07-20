package version1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hoangvuphone/internal/dtos"
	"hoangvuphone/internal/services"
)

type PermissionController struct {
	permissionService services.PermissionService
}

func NewPermissionController(permissionService services.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

// @Summary Get all permissions
// @Description Get all permissions
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dtos.PermissionResponse
// @Router /permissions [get]
func (pc *PermissionController) GetAllPermissions(c *gin.Context) {
	permissions, err := pc.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

// @Summary Get permission by ID
// @Description Get permission details by ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Success 200 {object} dtos.PermissionResponse
// @Router /permissions/{id} [get]
func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	permission, err := pc.permissionService.GetPermissionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permission)
}

// @Summary Create a new permission
// @Description Create a new permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dtos.CreatePermissionRequest true "Permission data"
// @Success 201 {object} dtos.PermissionResponse
// @Router /permissions [post]
func (pc *PermissionController) CreatePermission(c *gin.Context) {
	var input dtos.CreatePermissionRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission, err := pc.permissionService.CreatePermission(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, permission)
}

// @Summary Update a permission
// @Description Update permission details
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Param input body dtos.UpdatePermissionRequest true "Permission data"
// @Success 200 {object} dtos.PermissionResponse
// @Router /permissions/{id} [put]
func (pc *PermissionController) UpdatePermission(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input dtos.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission, err := pc.permissionService.UpdatePermission(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permission)
}

// @Summary Delete a permission
// @Description Delete a permission by ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Success 204
// @Router /permissions/{id} [delete]
func (pc *PermissionController) DeletePermission(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = pc.permissionService.DeletePermission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get permissions by role
// @Description Get all permissions assigned to a role
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param roleId path string true "Role ID"
// @Success 200 {array} dtos.PermissionResponse
// @Router /permissions/role/{roleId} [get]
func (pc *PermissionController) GetPermissionsByRole(c *gin.Context) {
	roleID, err := uuid.Parse(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}

	permissions, err := pc.permissionService.GetPermissionsByRole(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permissions)
}