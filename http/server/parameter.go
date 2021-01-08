package server

type (
	Request struct {
		Url     string `json:"url"`
		Content string `json:"content"`
	}
	
	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	
	DelayRequest struct {
		Request
	}
)
