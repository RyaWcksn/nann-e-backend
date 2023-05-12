package gpt

import "github.com/sashabaranov/go-openai"

type GptImpl struct {
	client *openai.Client
}

func NewGpt(APIKEY string) *GptImpl {
	client := openai.NewClient(APIKEY)
	return &GptImpl{
		client: client,
	}
}
