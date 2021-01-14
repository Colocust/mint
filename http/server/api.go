package server

import (
	"encoding/json"
	"mint/job"
	"mint/job/delay"
)

func Delay(content string, resp *Response) {
	req := new(DelayRequest)
	if err := json.Unmarshal([]byte(content), req); err != nil {
		resp.Code = StatusServerError
		return
	}

	node := new(job.Node)
	node.Method, node.Content, node.Url = req.Method, req.Content, req.Url

	queue := delay.GetInstance()
	queue.Push(req.When, node)

	resp.Code = StatusSuccess
	return
}
