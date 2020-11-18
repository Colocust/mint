package server

type Response struct {
	Code    int
	Message string
}

func NewResponse(code int, message string) *Response {
	return &Response{code, message}
}
