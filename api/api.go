package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"tinyQ/http/client"
)

func Ticker(writer http.ResponseWriter, r *http.Request) {
	type Request struct {
		Url      string                 `json:"url"`
		Data     map[string]interface{} `json:"data"`
		Method   string                 `json:"method"`
		Duration int64                  `json:"duration"`
	}
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	request := new(Request)
	json.Unmarshal(body, request)

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(request.Duration))

		select {
		case <-ticker.C:
			builder := new(client.NetBuilder)
			content, _ := json.Marshal(request.Data)
			resp, _ := builder.SetUrl(request.Url).SetContent(string(content)).NewNetSender().Send(request.Method)
			log.Println(resp)
		}
		ticker.Stop()
	}()

	response := Response{Code: http.StatusOK, Message: "OK"}
	res, _ := json.Marshal(response)
	writer.Write(res)
}
