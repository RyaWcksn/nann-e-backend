package ai

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math"
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
	fmt.Println(err)
	fmt.Println(err == sql.ErrNoRows)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		fmt.Println("Masuk sini pas looping")
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
		fmt.Println(chat)
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over result set: %v", err)
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
	fmt.Println("Keluar kesini ga")
	fmt.Println(chats)
	res := entities.Sessions{
		Id:        sessionId,
		CreatedAt: CreatedAt,
		LastChat:  chats[len(chats)-1].CreatedAt,
		Chats:     chats,
	}
	fmt.Println(res)
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
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}
	// Convert createdAtRaw ([]uint8) to time.Time
	createdAtStr := string(createdAtRaw)
	res.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
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
	var gender string
	err = a.DB.QueryRow("SELECT id, name, age, nanneId, gender FROM ai WHERE id = ? AND name = ?", payload.Id, payload.Name).Scan(&res.Id, &res.Name, &res.Age, &nanneId, &gender)
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
	var totalItems int
	err = a.DB.QueryRow("SELECT COUNT(*) FROM session where userId = ?", payload.Id).Scan(&totalItems)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err
	}

	if payload.Limit <= 0 {
		payload.Limit = 10
	}

	if payload.Page <= 0 {
		payload.Page = 1
	}

	offset := (payload.Page - 1) * payload.Limit
	totalPage := calculateTotalPages(totalItems, payload.Limit)

	sessionRow, err := a.DB.Query(`SELECT id, createdAt FROM session WHERE userId = ? ORDER BY createdAt DESC LIMIT ? OFFSET ?`, payload.Id, payload.Limit, offset)
	if err != nil {
		log.Printf("Error executing SQL query: %v", err)
		return nil, err
	}
	defer sessionRow.Close()

	var sessions []entities.Sessions
	for sessionRow.Next() {
		var session entities.Sessions
		var dateRaw []uint8
		err := sessionRow.Scan(&session.Id, &dateRaw)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		session.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(dateRaw))
		if err != nil {
			log.Printf("Error parsing createdAt: %v", err)
			return nil, err
		}
		sessions = append(sessions, session)
	}
	if err := sessionRow.Err(); err != nil {
		log.Printf("Error iterating over result set: %v", err)
		return nil, err
	}

	for i, _ := range sessions {
		chats, err := a.getAllChatsFromSessions(sessions[i].Id)
		if err != nil {
			return nil, fmt.Errorf("error getting chats from session: %v", err)
		}
		sessions[i].Chats = *chats
	}

	res.PageList = totalPage
	res.Gender = gender
	res.Sessions = sessions
	res.Nanne = nanneName

	return res, nil
}

func calculateTotalPages(totalItems, limit int) int {
	a := int(math.Ceil(float64(totalItems) / float64(limit)))
	return a
}

func (a AIImpl) getAllChatsFromSessions(session string) (res *[]entities.Chats, err error) {
	var resChat []entities.Chats
	rows, err := a.DB.Query(`select id, message, isUser, createdAt from chat where sessionId = ?`, session)
	if err != nil {
		log.Printf("Err := %v", err)
		return nil, err

	}
	for rows.Next() {
		var dateRaw []uint8
		var chat entities.Chats
		err = rows.Scan(&chat.Id, &chat.Message, &chat.IsUser, &dateRaw)
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err

		}
		chat.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(dateRaw))
		if err != nil {
			log.Printf("Err := %v", err)
			return nil, err
		}
		resChat = append(resChat, chat)
	}
	return &resChat, nil
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
