package routes

import (
	"hoangvuphone/internal/handlers/admin"

	"github.com/gin-gonic/gin"
)

// Đăng ký route cho admin dashboard
func RoleRoutes(r *gin.RouterGroup) {
	r.GET("/role", admin.RoleHandler)
}
