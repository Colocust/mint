package http

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
