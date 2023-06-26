package base

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func LoginHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
			},
		})
	}
}
