package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nann-e-backend/api/usecase"
	"github.com/nann-e-backend/config"
	"github.com/nann-e-backend/dtos"
	"github.com/nann-e-backend/entities"
)

type IHandler interface {
	Register(w http.ResponseWriter, r *http.Request) error
	Chat(w http.ResponseWriter, r *http.Request) error
	GetData(w http.ResponseWriter, r *http.Request) error
	GenerateUrl(w http.ResponseWriter, r *http.Request) error
	GenerateSession(w http.ResponseWriter, r *http.Request) error
	GetSession(w http.ResponseWriter, r *http.Request) error
	GetChatBySession(w http.ResponseWriter, r *http.Request) error
}

type ChatSessionPayload struct {
	SessionId string `json:"sessionId"`
}

type SessionPayload struct {
	UserId int `json:"userId"`
}

type GetSessionPayload struct {
	UserId int `json:"userId"`
}

type Handler struct {
	UC  usecase.IUsecase
	cfg config.Config
}

func (handler *Handler) GetChatBySession(w http.ResponseWriter, r *http.Request) error {
	var payload ChatSessionPayload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}
	chats, err := handler.UC.GetSessionBySessionId(payload.SessionId)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	resp := entities.Sessions{
		Id:        chats.Id,
		CreatedAt: chats.CreatedAt,
		LastChat:  chats.LastChat,
		Chats:     chats.Chats,
	}
	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (handler *Handler) GetSession(w http.ResponseWriter, r *http.Request) error {
	var payload GetSessionPayload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	sessions, err := handler.UC.GetSession(payload.UserId)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	var resp dtos.SessionResponse
	resp.Sessions = make([]dtos.Sessions, len(*sessions))
	for i, session := range *sessions {
		resp.Sessions[i] = dtos.Sessions{
			Id:        session.Id,
			CreatedAt: session.CreatedAt,
			LastChat:  session.LastChat,
		}
		resp.Sessions[i].Chats = make([]dtos.Chats, len(session.Chats))
		for _, chat := range session.Chats {
			resp.Sessions[i].Chats[i] = dtos.Chats{
				Id:        chat.Id,
				Message:   chat.Message,
				IsUser:    chat.IsUser,
				CreatedAt: chat.CreatedAt,
			}
		}
	}

	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (h Handler) GenerateUrl(w http.ResponseWriter, r *http.Request) (_ error) {
	var payload dtos.GenerateLinkRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	data, err := h.UC.GenerateUrl(payload)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	url := h.cfg.App.WEB + "?ref=" + data.Link
	resp := dtos.GenerateLinkResponse{
		Link: url,
	}

	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)

}

func NewHandler(u usecase.IUsecase, cfg config.Config) *Handler {
	return &Handler{
		UC:  u,
		cfg: cfg,
	}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload dtos.RegisterRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	data, err := h.UC.Register(payload)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	resp := dtos.RegisterResponse{
		Id:   data.Id,
		Name: data.Name,
	}

	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (h Handler) Chat(w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	payload := dtos.ChatRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	data, err := h.UC.Chat(payload)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	resp := dtos.ChatResponse{
		Response: data.Response,
	}
	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (h Handler) GetData(w http.ResponseWriter, r *http.Request) error {
	var payload dtos.DashboardParameter

	hash := r.URL.Query().Get("hash")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageint, _ := strconv.Atoi(page)
	limitint, _ := strconv.Atoi(limit)

	payload.Hash = hash
	payload.Limit = limitint
	payload.Page = pageint

	data, err := h.UC.GetData(payload)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}

	resp := entities.GetAiDatas{
		Id:       data.Id,
		Name:     data.Name,
		Age:      data.Age,
		Nanne:    data.Nanne,
		Gender:   data.Gender,
		PageList: data.PageList,
		Sessions: data.Sessions,
	}

	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (handler *Handler) GenerateSession(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	payload := SessionPayload{}
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Err := %v", err)
		return err
	}

	uuid, err := handler.UC.GenerateSession(payload.UserId)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	resp := dtos.GenerateSessionResponse{
		SessionId: uuid,
	}
	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}
