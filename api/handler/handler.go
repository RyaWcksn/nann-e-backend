package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nann-e-backend/api/usecase"
	"github.com/nann-e-backend/dtos"
)

type IHandler interface {
	Register(w http.ResponseWriter, r *http.Request) error
}

type Handler struct {
	UC usecase.IUsecase
}

func NewHandler(u usecase.IUsecase) *Handler {
	return &Handler{
		UC: u,
	}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) error {
	payload := dtos.RegisterRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("err := %v", err)
		return err
	}
	if err := json.Unmarshal(body, &payload); err != nil {
	}
	data, err := h.UC.Register(payload)
	if err != nil {
		log.Fatalf("err := %v", err)
		return err
	}
	resp := dtos.RegisterResponse{
		Id:   data.Id,
		Name: data.Name,
	}
	w.Header().Set("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(resp)
}
