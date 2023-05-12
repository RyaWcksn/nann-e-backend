package usecase

import (
	"log"

	"github.com/nann-e-backend/api/interfaces"
	"github.com/nann-e-backend/dtos"
	"github.com/nann-e-backend/entities"
)

type IUsecase interface {
	Register(r dtos.RegisterRequest) (resp *dtos.RegisterResponse, err error)
}

type UseCase struct {
	AI interfaces.IAi
}

func NewUsecase(a interfaces.IAi) *UseCase {
	return &UseCase{
		AI: a,
	}
}

func (u UseCase) Register(r dtos.RegisterRequest) (resp *dtos.RegisterResponse, err error) {
	payload := entities.RegisterEntity{
		Request: r,
	}
	data, err := u.AI.Register(payload)
	if err != nil {
		log.Fatalf("err := %v", err)
		return nil, err
	}
	resp = &dtos.RegisterResponse{
		Id:   data.Id,
		Name: data.Name,
	}
	return resp, nil
}
