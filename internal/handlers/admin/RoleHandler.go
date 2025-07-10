package admin

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/render"
)

func RoleHandler(c *gin.Context) {
	// Render the admin dashboard page
	render.RenderAdmin(c, "role", gin.H{
		"title": "Role Management",
		
	})
}