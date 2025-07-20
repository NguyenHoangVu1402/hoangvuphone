package routes

import (
	"hoangvuphone/internal/controllers/admin/version1"
	"github.com/gin-gonic/gin"

	"hoangvuphone/internal/services"
	"hoangvuphone/internal/middlewares"
)

// Đăng ký route cho admin dashboard
func RoleRoutes(router *gin.Engine, roleService services.RoleService, permissionService services.PermissionService) {
	// Đăng ký route cho trang dashboard
	

	roleController := version1.NewRoleController(roleService)
	router.GET("/admin/role", roleController.IndexRole)

	roleGroup := router.Group("/admin/v1/roles")
	{
		roleGroup.GET("", roleController.GetAllRoles)
		roleGroup.GET("/search", roleController.SearchRoles)
		roleGroup.POST("", middlewares.PermissionMiddleware("role.create", roleService), roleController.CreateRole)
		
		roleGroup.GET("/:id", roleController.GetRoleByID)
		roleGroup.PUT("/:id", middlewares.PermissionMiddleware("role.update", roleService), roleController.UpdateRole)
		roleGroup.DELETE("/:id", middlewares.PermissionMiddleware("role.delete", roleService), roleController.DeleteRole)
	}
}
