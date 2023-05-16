package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nann-e-backend/api/interfaces"
	"github.com/nann-e-backend/dtos"
	"github.com/nann-e-backend/entities"
	"github.com/nann-e-backend/pkgs/utils"
	"github.com/sashabaranov/go-openai"
)

type IUsecase interface {
	Register(r dtos.RegisterRequest) (resp *dtos.RegisterResponse, err error)
	Chat(r dtos.ChatRequest) (resp *dtos.ChatResponse, err error)
	GetData(r dtos.DashboardParameter) (resp *entities.GetAiDatas, err error)
	GenerateUrl(r dtos.GenerateLinkRequest) (resp *dtos.GenerateLinkResponse, err error)
}

type UseCase struct {
	AI  interfaces.IAi
	GPT interfaces.IGpt
}



func NewUsecase(a interfaces.IAi, g interfaces.IGpt) *UseCase {
	return &UseCase{
		AI:  a,
		GPT: g,
	}
}

func (u UseCase) Register(r dtos.RegisterRequest) (resp *dtos.RegisterResponse, err error) {
	payload := entities.RegisterEntity{
		Request: r,
	}
	data, err := u.AI.Register(payload)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	resp = &dtos.RegisterResponse{
		Id:   data.Id,
		Name: data.Name,
	}
	return resp, nil
}

func (u UseCase) Chat(r dtos.ChatRequest) (resp *dtos.ChatResponse, err error) {
	data, err := u.AI.GetData(r.Id, r.Name)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	oldChat, err := u.AI.GetChat(r.Id, r.Name)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	initiateContent := fmt.Sprintf(`
I want you to take the role as %s your name will be %s. 
You have to following this set of rules:

You will only talk something related to %s topic.
Please provide a simple response and easily can be understand by the %s year old. 
Introduce yourself to the kid and tell him/her what you will be helping with. Also provide me 2 simple and
related follow up question based on the topic or your previous response and 2 simple random question that can
continue the conversation with the kid. After that give me a random short fact that related to %s topic and can
easily understand by the %s years old kid.

Whenever you are prompted to provide a reply, always provide me a response as the following template:
Response : [ChatGPT as %s response]

Related Follow up question : [2 related question, in numbered format]

Random question : [2 random question, in numbered format]

Random fact : [Short fact that related to %s topic and can easily understand by the %s years old kid]

By using old chat like this %s

Generate follow up chat with this message %s
`, data.Role,
		data.Name,
		data.Role,
		data.Age,
		data.Role,
		data.Role,
		data.Age,
		oldChat.Chat,
		r.Message,
	)

	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: initiateContent,
		},
	}

	fmt.Println("Masuk sini ga")
	res, err := u.GPT.GenerateChat(message)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	fmt.Println("Masuk sini ga")
	followUpChat := fmt.Sprintf(`from this chat %s
please bautify the text, paragraph or anything that %s years old kid can understand
`, res.Choices[0].Message.Content, data.Age)

	message = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: followUpChat,
		},
	}
	folChat, err := u.GPT.FollowUpChat(message)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	fmt.Println(folChat.Choices[0].Message.Content)

	newChat := fmt.Sprintf("%s, %s", oldChat.Chat, folChat.Choices[0].Message.Content)
	err = u.AI.SaveChat(r.Id, r.Name, newChat)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	resp = &dtos.ChatResponse{
		Response: folChat.Choices[0].Message.Content,
	}
	return resp, err

}

func (u UseCase) GetData(r dtos.DashboardParameter) (resp *entities.GetAiDatas, err error) {

	var payload dtos.DashboardPayload
	decrypt := utils.Decrypt(r.Hash)
	fmt.Println(decrypt)
	err = json.Unmarshal([]byte(decrypt), &payload)
	if err != nil {
		log.Printf("Err := %v", err.Error())
		return nil, err
	}
	data, err := u.AI.GetAiDatas(payload)
	if err != nil {
		log.Printf("Err := %v", err.Error())
		return nil, err
	}

	resp = &entities.GetAiDatas{
		Name: data.Name,
		Age: data.Age,
		Role: data.Role,
		Chat: data.Chat,
	}

	return resp, nil
}

func (u UseCase) GenerateUrl(r dtos.GenerateLinkRequest) (resp *dtos.GenerateLinkResponse, err error) {
	generateJson := fmt.Sprintf(`{"id":%d,"name":"%s"}`, r.Id, r.Name)
	url := utils.Encrypt(generateJson)
	resp = &dtos.GenerateLinkResponse{
		Link: url,
	}
	return resp, nil
}
