package routes

import (
	"hoangvuphone/internal/handlers/admin"

	"github.com/gin-gonic/gin"
)

// Đăng ký route cho admin dashboard
func DashboardRoutes(r *gin.RouterGroup) {
	r.GET("/dashboard", admin.AdminDashboard)
}
