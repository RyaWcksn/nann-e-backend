package ai

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

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

func (aiimpl *AIImpl) GetChatBySessionId(sessionId string) (resp *entities.Sessions, err error) {
	var chats []entities.Chats
	rows, err := aiimpl.DB.Query(`SELECT id, message, isUser, createdAt FROM chat where sessionId = ?`, sessionId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	for rows.Next() {
		var chat entities.Chats
		var createdAtRaw []uint8
		err = rows.Scan(&chat.Id, &chat.Message, &chat.IsUser, &createdAtRaw)
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err
		}
		chat.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err
		}
		chats = append(chats, chat)
	}
	var sCATR []uint8
	err = aiimpl.DB.QueryRow(`select createdAt from session where id = ?`, sessionId).Scan(&sCATR)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	CreatedAt, err := time.Parse("2006-01-02 15:04:05", string(sCATR))
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	res := entities.Sessions{
		Id: sessionId,
		CreatedAt: CreatedAt,
		LastChat: chats[len(chats)-1].CreatedAt,
		Chats: chats,
	}
	return &res, nil

}

func (aiimpl *AIImpl) GetSession(userId int) (resp *[]entities.Sessions, err error) {
	var sessions []entities.Sessions

	rows, err := aiimpl.DB.Query(`SELECT id, createdAt FROM session WHERE userId = ?`, userId)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	for rows.Next() {
		var session entities.Sessions
		var sessionCreatedRaw []uint8
		err = rows.Scan(&session.Id, &sessionCreatedRaw)
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err
		}

		session.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(sessionCreatedRaw))
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err
		}

		sessions = append(sessions, session)
	}
	return &sessions, nil
}

func randString() (string, error) {
	// Generate a random byte slice of 16 bytes (128 bits)
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Failed to generate random string:", err)
		return "", nil
	}

	// Convert the byte slice to a hex string
	randomString := hex.EncodeToString(randomBytes)

	// Insert dashes every 8 characters
	finalString := randomString[:8] + "-" + randomString[8:16] + "-" + randomString[16:24] + "-" + randomString[24:]
	return finalString, nil

}

func (a *AIImpl) CreateSession(userId int) (id string, err error) {
	stmt, err := a.DB.Prepare("INSERT INTO session (id, userId) VALUES (?, ?)")
	if err != nil {
		log.Printf("Err := %v", err)
		return "", err
	}
	uuid, err := randString()
	if err != nil {
		log.Printf("Err := %v", err)
		return "", err
	}
	_, err = stmt.Exec(uuid, userId)
	if err != nil {
		log.Printf("Err := %v", err)
		return "", err
	}
	return uuid, nil
}

func (a AIImpl) Register(r entities.RegisterEntity) (resp *entities.RegisterEntityResponse, err error) {
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

func (a AIImpl) GetChat(id int, name string, isUser string, sessionId string) (resp *entities.GetChatEntityResp, err error) {
	res := &entities.GetChatEntityResp{}
	var createdAtRaw []uint8
	err = a.DB.QueryRow("SELECT message, createdAt FROM chat WHERE userId = ? AND isUser = ? AND sessionId = ? OR sessionId IS NULL ORDER BY createdAt DESC", id, isUser, sessionId).Scan(&res.Chat, &createdAtRaw)
	fmt.Println("Masuk")
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	// Convert createdAtRaw ([]uint8) to time.Time
	createdAtStr := string(createdAtRaw)
	res.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	fmt.Println("Masuk gasie")
	fmt.Println(res)
	return res, nil
}

// Save chat.
func (a AIImpl) SaveChat(id int, nanneId int, name string, chat string, isUser string, sessionId string) error {
	update, err := a.DB.Exec("INSERT into chat (userId, nanneId, message, isUser, sessionId) values (?, ?, ?, ?, ?)", id, nanneId, chat, isUser, sessionId)
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
