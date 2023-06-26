package quiz

import (
	"context"
	"log"
	"strconv"

	"github.com/boskeyacht/museapi/db/models"
	"github.com/boskeyacht/museapi/internal/types"
	"github.com/boskeyacht/museapi/internal/utils"
	"github.com/boskeyacht/museapi/prompts"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
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

func NewQuizHandler(ctx context.Context, db *bun.DB, openai *openai.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		generatePrompt := prompts.DefaultGenerateQuizRequest()
		generateRes := prompts.DefaultGenerateQuizResponse()
		quiz := models.DefaultQuiz()

		generatePrompt.FillAttributes(
			types.NewAttribute("{{topic}}", ""),
			types.NewAttribute("{{examples}}", ""),
			types.NewAttribute("{{question_limit}}", ""),
			types.NewAttribute("{{difficulty}}", ""),
		)

		err := utils.SendAndUnmarshal(ctx, openai, generatePrompt.Prompt, generateRes)
		if err != nil {
			c.JSON(400, gin.H{
				"error":   "Bad request",
				"message": "failed to generate quiz",
			})

			return
		}

		log.Printf("Generated quiz: %v\n", generateRes)

		quiz.Title = generateRes.Topic

		for _, q := range generateRes.Questions {
			question := models.DefaultQuestion()

			question.Question = q.Question
			question.CorrectAnswer = q.Answer
			question.PossibleAnswers = q.PossibleAnswers

			switch q.Type {
			case "MultipleChoiceQuestion":
				question.Type = types.MultipleChoiceQuestion

			case "ShortAnswerQuestion":
				question.Type = types.ShortAnswerQuestion

			case "TrueFalseQuestion":
				question.Type = types.TrueFalseQuestion

			default:
				question.Type = types.MultipleChoiceQuestion
			}

			quiz.Questions = append(quiz.Questions, question)
		}

		err = quiz.NewQuiz(ctx, db)
		if err != nil {
			c.JSON(400, gin.H{
				"error":   "Bad request",
				"message": "failed to create quiz",
			})

			return
		}

		for _, q := range quiz.Questions {
			q.QuizID = quiz.ID

			err = q.NewQuestion(ctx, db)
			if err != nil {
				c.JSON(400, gin.H{
					"error":   "Bad request",
					"message": "failed to create question",
				})

				return
			}
		}

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"quiz": quiz,
			},
		})
	}
}

func ScoreQuizHandler(ctx context.Context, db *bun.DB, openai *openai.Client) gin.HandlerFunc {
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

		for _, q := range quiz.Questions {
			if q.Type != types.ShortAnswerQuestion {
				if q.UserAnswer == q.CorrectAnswer {
					quiz.Score++

					q.IsCorrect = true
				}
			} else {
				// @todo score short answer questions
				continue
			}

		}

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"id":    quiz.ID,
				"score": quiz.Score,
			},
		})
	}
}

func UpdateQuizHandler(ctx context.Context, db *bun.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string
		quiz := models.DefaultQuiz()

		if c.GetHeader("If-Match") == "" {
			c.JSON(400, gin.H{
				"error":    "Bad request",
				"messsage": "If-Match header is required",
			})

			return
		} else {
			id = c.GetHeader("If-Match")
		}

		err := c.BindJSON(&quiz)
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

		quiz.ID = int64(idc)

		err = quiz.UpdateQuiz(ctx, db)
		if err != nil {
			c.JSON(500, gin.H{
				"error":    "Internal server error",
				"messsage": "failed to update quiz",
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
