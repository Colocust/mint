package server

type (
	Request struct {
		Url     string `json:"url"`
		Content string `json:"content"`
	}

	Response struct {
		Code    int    `json:"code"`
	}

	DelayRequest struct {
		Request
	}
)