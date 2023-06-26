package prompts

import (
	"strings"

	"github.com/boskeyacht/museapi/internal/types"
)

type GenerateQuizRequest struct {
	Prompt string `json:"prompt"`
}

func DefaultGenerateQuizRequest() *GenerateQuizRequest {
	return &GenerateQuizRequest{
		Prompt: `You are a {{topic}} expert and I am a student. You are evaluating my understanding of {{topic}} by asking me to questions in the form of a "quiz". This means you must provide the question along with the correct answer.
		Here are some guidelines:

		Examples: {{examples}}
		Question Limit: {{question_limit}}
		Difficulty: {{difficulty}}
		
		1. Your questions should elicit medium-sized explanations as responses.
		2. Your questions should be technical in nature and assess my knowledge.
		3. Your questions should primarily focus on practical details rather than theory.
		
		Use the following JSON schema for your response. Do not return anything except for the JSON object
		{
			"topic": " Quiz Topic",
			"difficulty": "easy",
			"questions": [
				{
					"question": "Question 1",
					"answer": "Answer 1"
					"possible_answers": ["Answer 1", "Answer 2", "Answer 3"], # if applicable, else empty array
					"type": "multiple_choice" # or short answer, or true/false
				},
			],
		}`,
	}
}

func (gqr *GenerateQuizRequest) FillAttributes(attrs ...*types.Attribute) {
	for _, attr := range attrs {
		gqr.Prompt = strings.Replace(gqr.Prompt, attr.Original, attr.Replacement, -1)
	}
}

type GenerateQuizResponse struct {
	Topic      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Questions  []*struct {
		Type            string   `json:"type"`
		Question        string   `json:"question"`
		Answer          string   `json:"answer"`
		PossibleAnswers []string `json:"possible_answers"`
	} `json:"questions"`
}

func DefaultGenerateQuizResponse() *GenerateQuizResponse {
	return &GenerateQuizResponse{
		Topic:      "",
		Difficulty: "",
		Questions: []*struct {
			Type            string   `json:"type"`
			Question        string   `json:"question"`
			Answer          string   `json:"answer"`
			PossibleAnswers []string `json:"possible_answers"`
		}{},
	}
}
