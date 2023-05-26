package dtos

import "time"

// ffjson: RegisterRequest
type RegisterRequest struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     string `json:"age"`
	NanneId int    `json:"nanneId"`
}

type RegisterResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ChatRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SessionId string `json:"sessionId"`
	Message   string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type DashboardParameter struct {
	Hash string `json:"hash"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

type DashboardPayload struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Page int    `json:"page"`
}

type GenerateLinkRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GenerateLinkResponse struct {
	Link string `json:"link"`
}

type GenerateSessionResponse struct {
	SessionId string `json:"sessionId"`
}

type SessionResponse struct {
	Sessions []Sessions `json:"sessions"`
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
