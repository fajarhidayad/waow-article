package common

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
