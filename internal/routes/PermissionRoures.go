package routes

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/controllers/admin/version1"
	"hoangvuphone/internal/middlewares"
	"hoangvuphone/internal/services"
)

func PermissionRoutes(router *gin.Engine, roleService services.RoleService, permissionService services.PermissionService) {
	permissionController := version1.NewPermissionController(permissionService)

	permGroup := router.Group("/permissions")
	{
		permGroup.GET("", middlewares.PermissionMiddleware("permission.view", roleService), permissionController.GetAllPermissions)
		permGroup.POST("", middlewares.PermissionMiddleware("permission.create", roleService), permissionController.CreatePermission)
		
		permGroup.GET("/:id", middlewares.PermissionMiddleware("permission.view", roleService), permissionController.GetPermissionByID)
		permGroup.PUT("/:id", middlewares.PermissionMiddleware("permission.update", roleService), permissionController.UpdatePermission)
		permGroup.DELETE("/:id", middlewares.PermissionMiddleware("permission.delete", roleService), permissionController.DeletePermission)
		
		permGroup.GET("/role/:roleId", middlewares.PermissionMiddleware("permission.view", roleService), permissionController.GetPermissionsByRole)
	}
}