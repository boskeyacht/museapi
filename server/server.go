package server

import (
	"context"
	"sync"

	"github.com/boskeyacht/museapi/db/models"
	"github.com/boskeyacht/museapi/internal/types"
	"github.com/boskeyacht/museapi/server/handlers/base"
	"github.com/boskeyacht/museapi/server/handlers/flashcards"
	"github.com/boskeyacht/museapi/server/handlers/quiz"
	"github.com/boskeyacht/museapi/server/handlers/user"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/uptrace/bun"
)

type Server struct {
	DB     *bun.DB
	OpenAI *openai.Client
	Config *types.Config
	sync.Mutex
}

func NewServer(cfg *types.Config, db *bun.DB, openai *openai.Client) *Server {
	return &Server{
		DB:     db,
		Config: cfg,
		OpenAI: openai,
	}
}

func (s *Server) Close() error {
	return s.DB.Close()
}

func (s *Server) InitRoutes(ctx context.Context) *gin.Engine {
	server := gin.Default()

	baseRoute := server.Group("/")
	{
		baseRoute.POST("login", base.LoginHandler(ctx, s.DB))
		baseRoute.POST("signup", base.SignUpHandler(ctx, s.DB))
	}

	userRoute := server.Group("/user")
	{
		userRoute.GET("/:id", user.GetUserHandler(ctx, s.DB))
		userRoute.PATCH("/:id", user.UpdateUserHandler(ctx, s.DB))
	}

	quizRoute := server.Group("/quiz")
	{
		quizRoute.GET("/:id", quiz.GetQuizHandler(ctx, s.DB))
		quizRoute.POST("", quiz.NewQuizHandler(ctx, s.DB, s.OpenAI))
		quizRoute.POST("/:id/score", quiz.ScoreQuizHandler(ctx, s.DB, s.OpenAI))
		quizRoute.PATCH("/", quiz.UpdateQuizHandler(ctx, s.DB))
	}

	chatRoute := server.Group("/chat")
	{
		chatRoute.GET("/:id", quiz.GetQuizHandler(ctx, s.DB))
		chatRoute.POST("", quiz.NewQuizHandler(ctx, s.DB, s.OpenAI))
		chatRoute.PATCH("/:id", quiz.UpdateQuizHandler(ctx, s.DB))
	}

	flashcardsRoute := server.Group("/flashcards")
	{
		flashcardsRoute.POST("", flashcards.NewFlashcardsHandler(ctx, s.DB))
	}

	return server
}

func (s *Server) InitTables(ctx context.Context) error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Quiz)(nil),
		(*models.Flashcard)(nil),
	}

	for _, model := range models {
		if _, err := s.DB.NewCreateTable().Model(model).Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}
