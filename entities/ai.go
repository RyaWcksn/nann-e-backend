package entities

import "github.com/nann-e-backend/dtos"

type RegisterEntity struct {
	Request dtos.RegisterRequest
}

type RegisterEntityResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
