package gpt

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type GptImpl struct {
	client *openai.Client
}

func NewGpt(APIKEY string) *GptImpl {
	client := openai.NewClient(APIKEY)
	return &GptImpl{
		client: client,
	}
}

func (g *GptImpl) GenerateChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error) {
	resp, err := g.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: message,
		},
	)
	if err != nil {
		log.Printf("Err := %v from OPENAI", err)
		return nil, err
	}
	return &resp, err
}

func (g *GptImpl) FollowUpChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error) {
	resp, err := g.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: message,
		},
	)
	if err != nil {
		log.Printf("Err := %v from OPENAI", err)
		return nil, err
	}

	return &resp, err
}
