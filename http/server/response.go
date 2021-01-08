package server

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(code int, message string) *Response {
	return &Response{code, message}
}
