package handlers

import "github.com/gin-gonic/gin"

type BaseHandler struct{}

func (b *BaseHandler) GetUserIdFromContext(c *gin.Context) (int64, bool) {
	userId, exists := c.Get("userId")
	if !exists {
		return 0, false
	}

	// Type assertion mit error handling
	if id, ok := userId.(int64); ok {
		return id, true
	}
	return 0, false
}
