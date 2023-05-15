package interfaces

import (
	"github.com/nann-e-backend/entities"
	"github.com/sashabaranov/go-openai"
)

type IAi interface {
	Register(r entities.RegisterEntity) (resp *entities.RegisterEntityResponse, err error)
	GetData(id int, name string) (resp *entities.GetDataEntityResp, err error)
	GetChat(id int, name string) (resp *entities.GetChatEntityResp, err error)
	SaveChat(id int, name string, chat string) error
}

type IGpt interface {
	GenerateChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error)
	FollowUpChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error)
}
