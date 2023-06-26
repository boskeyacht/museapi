package user

import (
	"context"
	"strconv"

	"github.com/boskeyacht/museapi/db/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetUserHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string

		if c.Params.ByName("id") == "" {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "id is required",
			})

			return

		} else {
			id = c.Params.ByName("id")
		}

		user := models.DefaultUser()
		idc, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "id is not a number",
			})

			return
		}

		user.ID = int64(idc)

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"user": user,
			},
		})
	}
}

func UpdateUserHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string
		user := models.DefaultUser()

		if c.GetHeader("If-Match") == "" {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "If-Match header is required",
			})

			return
		} else {
			id = c.GetHeader("If-Match")
		}

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "failed to bind json",
			})

			return
		}

		idc, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "id is not a number",
			})

			return
		}

		user.ID = int64(idc)

		err = user.UpdateUser(ctx, db)
		if err != nil {
			c.JSON(500, gin.H{
				"error":    "Internal server error",
				"messsage": "failed to update user",
			})

			return
		}

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"user": user,
			},
		})
	}
}
