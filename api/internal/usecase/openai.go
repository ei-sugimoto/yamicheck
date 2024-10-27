package usecase

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIServiceImp struct {
	Client *openai.Client
}

func NewOpenAIService(cilent *openai.Client) *OpenAIServiceImp {
	return &OpenAIServiceImp{
		Client: cilent,
	}
}

func (s *OpenAIServiceImp) GenerateText() (string, error) {
	resp, err := s.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "こんにちは",
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
