package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/models"
	"hoangvuphone/internal/services"
)

func PermissionMiddleware(permissionSlug string, roleService services.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user from context (set by auth middleware)
		user, exists := c.Get("currentUser")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		account, ok := user.(models.Account)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			return
		}

		// Check if user has the required permission
		hasPermission := false
		for _, role := range account.Roles {
			permissions, err := roleService.GetPermissions(role.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check permissions"})
				return
			}

			for _, perm := range permissions {
				if perm.Slug == permissionSlug {
					hasPermission = true
					break
				}
			}

			if hasPermission {
				break
			}
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		c.Next()
	}
}