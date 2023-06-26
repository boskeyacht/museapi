package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

func SendAndUnmarshal(ctx context.Context, client openai.Client, prompt string, respTarget interface{}) (interface{}, error) {
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)

		return nil, err
	}

	fmt.Println(resp.Choices[0].Message.Content)

	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), respTarget)
	if err != nil {
		log.Printf("Unmarshal error: %v\n", err)

		return nil, err
	}

	return respTarget, nil
}
