package ai

import (
	"database/sql"
	"log"

	"github.com/nann-e-backend/dtos"
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
	stmt, err := a.DB.Prepare("INSERT INTO ai (name, role, age, chat) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	query, err := stmt.Exec(r.Request.Name, r.Request.Role, r.Request.Age, "Initial message")
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	id, err := query.LastInsertId()
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	resp = &entities.RegisterEntityResponse{
		Id:   int(id),
		Name: r.Request.Name,
	}

	return resp, nil
}

func (a AIImpl) GetData(id int, name string) (resp *entities.GetDataEntityResp, err error) {
	res := &entities.GetDataEntityResp{}
	err = a.DB.QueryRow("SELECT name, age, role FROM ai WHERE id = ? AND name = ?", id, name).Scan(&res.Name, &res.Age, &res.Role)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	return res, nil
}

func (a AIImpl) GetChat(id int, name string) (resp *entities.GetChatEntityResp, err error) {
	res := &entities.GetChatEntityResp{}
	err = a.DB.QueryRow("SELECT chat FROM ai WHERE id = ? AND name = ?", id, name).Scan(&res.Chat)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	return res, nil
}

// Save chat.
func (a AIImpl) SaveChat(id int, name string, chat string) error {
	update, err := a.DB.Exec("UPDATE ai SET chat = ? WHERE id = ? AND name = ?", chat, id, name)
	if err != nil {
		log.Printf("Err := %v", err)
		return err
	}
	if _, err := update.LastInsertId(); err != nil {
		log.Printf("Err := %v", err)
		return err
	}

	return nil
}

func (a AIImpl) GetAiDatas(payload dtos.DashboardPayload) (resp *entities.GetAiDatas, err error) {
	res := &entities.GetAiDatas{}
	err = a.DB.QueryRow("SELECT name, age, role, chat FROM ai WHERE id = ? AND name = ?", payload.Id, payload.Name).Scan(&res.Name, &res.Age, &res.Role, &res.Chat)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	return res, nil
}
