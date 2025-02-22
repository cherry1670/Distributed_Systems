package models

type Response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type Operations interface {
	ProcessCreate(Task) Response
}
