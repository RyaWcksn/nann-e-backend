package interfaces

import (
	"github.com/nann-e-backend/dtos"
	"github.com/nann-e-backend/entities"
	"github.com/sashabaranov/go-openai"
)

type IAi interface {
	Register(r entities.RegisterEntity) (resp *entities.RegisterEntityResponse, err error)
	GetData(id int, name string) (resp *entities.GetDataEntityResp, err error)
	GetChat(id int, name string, isUser string, sessionId string) (resp *entities.GetChatEntityResp, err error)
	SaveChat(id int, nanneId int, name string, chat string, isUser string, sessionId string) error
	GetAiDatas(payload dtos.DashboardPayload) (resp *entities.GetAiDatas, err error)
	GetAiInfo(nanneId int) (resp *entities.AiData, err error)
	CreateSession(userId int) (id string, err error)
	GetSession(userId int) (resp *[]entities.Sessions, err error)
	GetChatBySessionId(sessionId string) (resp *entities.Sessions, err error)
}

type IGpt interface {
	GenerateChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error)
	FollowUpChat(message []openai.ChatCompletionMessage) (res *openai.ChatCompletionResponse, err error)
}
