package entities

import (
	"time"

	"github.com/nann-e-backend/dtos"
)

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
	Chat      string
	CreatedAt time.Time
}

type GetAiDatas struct {
	Id       int
	Name     string
	Age      string
	Gender   string
	Nanne    string
	PageList int
	Sessions []Sessions
}

type Chat struct {
	Message   string
	IsUser    string
	CreatedAt time.Time
}

type AiData struct {
	Name        string
	Description string
}
type Sessions struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	LastChat  time.Time `json:"lastChat"`
	Chats     []Chats   `json:"chats"`
}

type Chats struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	IsUser    string    `json:"isUser"`
	CreatedAt time.Time `json:"createdAt"`
}
