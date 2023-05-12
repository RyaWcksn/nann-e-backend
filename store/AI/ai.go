package ai

import (
	"database/sql"
	"log"

	"github.com/nann-e-backend/entities"
)

type AIImpl struct {
	DB *sql.DB
}

func NewAi(DB *sql.DB) *AIImpl {
	return &AIImpl{
		DB: DB,
	}
}

func (a AIImpl) Register(r entities.RegisterEntity) (resp *entities.RegisterEntityResponse, err error) {
	stmt, err := a.DB.Prepare("INSERT INTO ai (name, role, age) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Err := %v", err)
		return nil, err
	}
	query, err := stmt.Exec(r.Request.Name, r.Request.Role, r.Request.Age)
	if err != nil {
		log.Fatalf("Err := %v", err)
		return nil, err
	}
	id, err := query.LastInsertId()
	if err != nil {
		log.Fatalf("Err := %v", err)
		return nil, err
	}
	resp = &entities.RegisterEntityResponse{
		Id:   int(id),
		Name: r.Request.Name,
	}

	return resp, nil
}
