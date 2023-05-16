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
}

type Handler struct {
	UC  usecase.IUsecase
	cfg config.Config
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
	id := r.Header.Get("id")
	name := r.Header.Get("name")
	payload := dtos.ChatRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	if err := json.Unmarshal(body, &payload); err != nil {
	}
	dataID, _ := strconv.Atoi(id)
	payload.Id = dataID
	payload.Name = name
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	data, err := h.UC.GetData(payload)
	if err != nil {
		log.Printf("Err := %v", err)

		return err
	}

	resp := entities.GetAiDatas{
		Name: data.Name,
		Age:  data.Age,
		Role: data.Role,
		Chat: data.Chat,
	}

	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}
