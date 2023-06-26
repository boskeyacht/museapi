package base

import (
	"context"
	"log"
	"net/http"

	"github.com/boskeyacht/museapi/db/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type SignUpRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func DefaultSignUpRequest() *SignUpRequest {
	return &SignUpRequest{
		Username:  "",
		FirstName: "",
		LastName:  "",
		Email:     "",
		Password:  "",
	}
}

// @todo handle username conflict
func SignUpHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.DefaultUser()
		sur := DefaultSignUpRequest()

		err := c.BindJSON(&sur)
		if err != nil {
			log.Printf("failed to bind user json: %v", err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		user.Username = sur.Username
		user.FirstName = sur.FirstName
		user.LastName = sur.LastName
		user.Email = sur.Email

		err = user.NewUser(ctx, db)
		if err != nil {
			log.Printf("failed to create user: %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"username": "boskeyacht",
				"id":       user.ID,
			},
		})
	}
}
