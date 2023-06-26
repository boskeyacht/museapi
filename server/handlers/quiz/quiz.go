package quiz

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
			},
		})
	}
}

func NewQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
			},
		})
	}
}

func ScoreQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
			},
		})
	}
}

func UpdateQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
			},
		})
	}
}
