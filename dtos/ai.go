package dtos

type RegisterRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Age  string `json:"age"`
}

type RegisterResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ChatRequest struct {
	Id      int
	Name    string
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type DashboardParameter struct {
	Id   int
	Name string
}
