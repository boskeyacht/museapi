package quiz

import (
	"context"
	"strconv"

	"github.com/boskeyacht/museapi/db/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string

		quiz := models.DefaultQuiz()

		if c.Params.ByName("id") != "" {
			id = c.Params.ByName("id")
		} else {
			c.JSON(400, gin.H{
				"error":   "Bad request",
				"message": "id is required",
			})

			return
		}

		idc, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error":   "Bad request",
				"message": "id must be an integer",
			})

			return
		}

		quiz.ID = int64(idc)

		err = quiz.GetQuiz(ctx, db)
		if err != nil {
			c.JSON(400, gin.H{
				"error":   "Bad request",
				"message": "failed to get quiz",
			})

			return
		}

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"quiz": quiz,
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
