package ai

import (
	"database/sql"
	"fmt"
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
	fmt.Println(r)
	stmt, err := a.DB.Prepare("INSERT INTO ai (name, gender, age, nanneId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	query, err := stmt.Exec(r.Request.Name, r.Request.Gender, r.Request.Age, r.Request.NanneId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	id, err := query.LastInsertId()
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	stmt2, err := a.DB.Prepare("INSERT INTO chat (userId, nanneId, message, isUser) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	_, err = stmt2.Exec(id, r.Request.NanneId, "Hello, how can i help you today?", "no")
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
	fmt.Println(id, name)
	err = a.DB.QueryRow("SELECT name, age, gender, nanneId FROM ai WHERE id = ? AND name = ?", id, name).Scan(&res.Name, &res.Age, &res.Gender, &res.NanneId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	return res, nil
}

func (a AIImpl) GetChat(id int, name string, isUser string) (resp *entities.GetChatEntityResp, err error) {
	res := &entities.GetChatEntityResp{}
	err = a.DB.QueryRow("SELECT message, createdAt FROM chat WHERE userId = ? AND isUser = ? ORDER BY createdAt DESC", id, isUser).Scan(&res.Chat, &res.CreatedAt)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	return res, nil
}

// Save chat.
func (a AIImpl) SaveChat(id int, nanneId int, name string, chat string, isUser string) error {
	update, err := a.DB.Exec("INSERT into chat (userId, nanneId, message, isUser) values (?, ?, ?, ?)", id, nanneId, chat, isUser)
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
	var nanneId int
	err = a.DB.QueryRow("SELECT id, name, age, nanneId FROM ai WHERE id = ? AND name = ?", payload.Id, payload.Name).Scan(&res.Id, &res.Name, &res.Age, &nanneId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	var nanneName string
	err = a.DB.QueryRow("SELECT name FROM nanne WHERE id = ? ", nanneId).Scan(&nanneName)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err

	}
	err = a.DB.QueryRow("SELECT name FROM nanne WHERE id = ? ", nanneId).Scan(&nanneName)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err

	}
	var chats []entities.Chat
	rows, err := a.DB.Query("SELECT message, isUser FROM chat WHERE userId = ? AND nanneId = ?", payload.Id, nanneId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err

	}

	for rows.Next() {
		var chat entities.Chat
		err = rows.Scan(&chat.Message, &chat.IsUser)
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err

		}
		chats = append(chats, chat)
	}

	res.Nanne = nanneName
	res.Chat = chats

	return res, nil
}

func (a AIImpl) GetAiInfo(nanneId int) (resp *entities.AiData, err error) {
	res := &entities.AiData{}
	err = a.DB.QueryRow("SELECT name, description FROM nanne where id = ?", nanneId).Scan(&res.Name, &res.Description)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	return res, nil
}
