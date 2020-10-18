package task

import (
	"mint/http/server"
	"net/http"
)

func Add(req *AddTaskRequest) *AddTaskResponse {
	t := &Task{req.Url, req.Data}
	go product(t)

	resp := new(AddTaskResponse)
	resp.Code = http.StatusOK
	resp.Message = "OK"
	return resp
}

type (
	AddTaskRequest struct {
		Url  string                 `json:"url"`
		Data map[string]interface{} `json:"data"`
	}
	AddTaskResponse struct {
		server.Response
	}
)
