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
	Name string
	Age  string
	Role string
}

type GetChatEntityResp struct {
	Chat string
}

type GetAiDatas struct {
	Name string
	Age  string
	Role string
	Chat string
}
