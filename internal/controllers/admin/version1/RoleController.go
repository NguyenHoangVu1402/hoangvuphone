package version1

import (
	"hoangvuphone/internal/render"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hoangvuphone/internal/dtos"
	"hoangvuphone/internal/services"
	"hoangvuphone/internal/validations"
)



type RoleController struct{
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (rc *RoleController) IndexRole(c *gin.Context) {
    // Lấy danh sách roles từ service
    roles, _, err := rc.roleService.GetAllRoles("", 1, 10) 
    if err != nil {
        // Xử lý lỗi nếu cần
        render.RenderAdmin(c, "dashboard", gin.H{
            "title": "Dashboard",
            "error": "Failed to load roles data",
        })
        return
    }

    // Render template với dữ liệu roles
    render.RenderAdmin(c, "role", gin.H{
        "title": "Role Management",
        "roles": roles,
    })
}

// @Summary Get all roles
// @Description Get all roles with pagination and search
// @Tags Roles
// @Accept json
// @Produce json
// @Param search query string false "Search term"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} dtos.PaginatedRoleResponse
// @Router /roles [get]
func (rc *RoleController) GetAllRoles(c *gin.Context) {
	search := c.Query("search")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	roles, total, err := rc.roleService.GetAllRoles(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dtos.PaginatedRoleResponse{
		Total: total,
		Page:  page,
		Limit: limit,
		Data:  roles,
	}

	c.JSON(http.StatusOK, response)
}


// @Summary Get role by ID
// @Description Get role details by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} dtos.RoleResponse
// @Router /roles/{id} [get]
func (rc *RoleController) GetRoleByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	role, err := rc.roleService.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// @Summary Create a new role
// @Description Create a new role with permissions
// @Tags Roles
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dtos.CreateRoleRequest true "Role data"
// @Success 201 {object} dtos.RoleResponse
// @Router /roles [post]
func (rc *RoleController) CreateRole(c *gin.Context) {
	var input dtos.CreateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validations.ValidateCreateRole(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := rc.roleService.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// @Summary Update a role
// @Description Update role details and permissions
// @Tags Roles
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Role ID"
// @Param input body dtos.UpdateRoleRequest true "Role data"
// @Success 200 {object} dtos.RoleResponse
// @Router /roles/{id} [put]
func (rc *RoleController) UpdateRole(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input dtos.UpdateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validations.ValidateUpdateRole(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := rc.roleService.UpdateRole(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Summary Delete a role
// @Description Delete a role by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Role ID"
// @Success 204
// @Router /roles/{id} [delete]
func (rc *RoleController) DeleteRole(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = rc.roleService.DeleteRole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search roles
// @Description Search roles by name or slug with autocomplete
// @Tags Roles
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} dtos.RoleResponse
// @Router /roles/search [get]
func (rc *RoleController) SearchRoles(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	roles, err := rc.roleService.SearchRoles(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}