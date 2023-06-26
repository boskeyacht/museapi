package prompts

import (
	"strings"

	"github.com/boskeyacht/museapi/internal/types"
)

type GenerateQuizRequest struct {
	Prompt string `json:"prompt"`
}

func NewGenerateQuizRequest(data map[string]interface{}) *GenerateQuizRequest {
	return &GenerateQuizRequest{
		Prompt: "You are",
	}
}

func (gqr *GenerateQuizRequest) FillAttributes(attrs ...*types.Attribute) {
	for _, attr := range attrs {
		gqr.Prompt = strings.Replace(gqr.Prompt, attr.Original, attr.Replacement, -1)
	}
}

type GenerateQuizResponse struct {
	Title string `json:"title"`
}

func DefaultGenerateQuizResponse() *GenerateQuizResponse {
	return &GenerateQuizResponse{
		Title: "",
	}
}
