package server

type (
	Request struct {
		Url     string `json:"url"`
		Content string `json:"content"`
		Method  string `json:"method"`
	}

	Response struct {
		Code int `json:"code"`
	}

	DelayRequest struct {
		Request
		When int `json:"when"`
	}
)
