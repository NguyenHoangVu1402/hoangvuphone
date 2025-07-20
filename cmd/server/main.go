package main

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/config"
	"hoangvuphone/internal/render"
	"hoangvuphone/internal/routes"
	"hoangvuphone/internal/migrations"

	"hoangvuphone/internal/repositories"
	"hoangvuphone/internal/services"
)

func main() {
	// Kết nối database
	config.ConnectDatabase()

	migrations.MigrateDatabase(config.DB)

	// Tải giao diện templates
	render.LoadTemplates()

	// Khởi tạo Gin router
	router := gin.Default()

	// Load template + static file
	router.LoadHTMLGlob("web/templates/**/*.html")
	router.Static("/static", "./web/static")

	// 🔧 Khởi tạo repository
	roleRepo := repositories.NewRoleRepository(config.DB)
	permissionRepo := repositories.NewPermissionRepository(config.DB)

	// 🔧 Khởi tạo service
	roleService := services.NewRoleService(roleRepo, permissionRepo)
	permissionService := services.NewPermissionService(permissionRepo)


	// Nhóm route /admin
	routes.DashboardRoutes(router)
	routes.RoleRoutes(router, roleService, permissionService)
	routes.PermissionRoutes(router, roleService, permissionService)

	// Chạy server với port từ config
	router.Run(":" + config.GetPort())
}
