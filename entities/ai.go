package entities

import "github.com/nann-e-backend/dtos"

type RegisterEntity struct {
	Request dtos.RegisterRequest
}

type RegisterEntityResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetDataEntityResp struct {
	Name    string
	Age     string
	Gender  string
	NanneId int
}

type GetChatEntityResp struct {
	Chat string
}

type GetAiDatas struct {
	Id    int
	Name  string
	Age   string
	Nanne string
	Chat  []Chat
}

type Chat struct {
	Message string
	IsUser  string
}

type AiData struct {
	Name        string
	Description string
}
