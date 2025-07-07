package main

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/config"
	"hoangvuphone/internal/render"
	"hoangvuphone/internal/routes"
)

func main() {
	// Kết nối database
	config.ConnectDatabase()

	// Tải giao diện templates
	render.LoadTemplates()

	// Khởi tạo Gin router
	router := gin.Default()

	// Nạp template HTML và static
	router.LoadHTMLGlob("web/templates/**/*.html")
	router.Static("/static", "./web/static")

	// Nhóm route /admin
	adminGroup := router.Group("/admin")
	routes.DashboardRoutes(adminGroup)

	// Chạy server với port từ config
	router.Run(":" + config.GetPort())
}
