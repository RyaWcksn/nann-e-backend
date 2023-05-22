package dtos

// ffjson: RegisterRequest
type RegisterRequest struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     string `json:"age"`
	NanneId int    `json:"nanneId"`
}

type RegisterResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ChatRequest struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type DashboardParameter struct {
	Hash string `json:"hash"`
}

type DashboardPayload struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GenerateLinkRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GenerateLinkResponse struct {
	Link string `json:"link"`
}
