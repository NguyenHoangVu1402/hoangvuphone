package version1

import (
	"hoangvuphone/internal/render"
	"github.com/gin-gonic/gin"
)

func IndexDashboard(c *gin.Context) {
	// Render the admin dashboard page
	render.RenderAdmin(c, "dashboard", gin.H{
		"title": "Dashboard",
	})
}