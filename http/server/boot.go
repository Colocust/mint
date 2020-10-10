package server

import (
	"encoding/json"
	"net/http"
	"tinyQ/src/api"
	"tinyQ/src/parameter"
)

func Runner() {
	var ret parameter.Response
	req := new(parameter.Request)

	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	json.Unmarshal(body, req)


	api.Ticker(*req)
	ret = parameter.Response{Code: http.StatusOK, Message: "OK"}

	result, _ := json.Marshal(ret)
	_, _ = writer.Write(result)
}
