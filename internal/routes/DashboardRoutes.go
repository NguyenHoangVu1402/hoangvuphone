package routes

import (
	"hoangvuphone/internal/controllers/admin/version1"

	"github.com/gin-gonic/gin"
)

// Đăng ký route cho admin dashboard
func DashboardRoutes(router  *gin.Engine) {
	dashboardGroup := router.Group("/admin")
	dashboardGroup.GET("/dashboard", version1.IndexDashboard)
}
