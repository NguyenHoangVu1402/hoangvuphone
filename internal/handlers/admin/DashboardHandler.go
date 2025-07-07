package admin

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/render"
)

func AdminDashboard(c *gin.Context) {
	// Render the admin dashboard page
	render.RenderAdmin(c, "dashboard", gin.H{
		"title": "Dashboard",
		
	})
}