package api

import (
	"encoding/json"
	"log"
	"time"
	"tiny-tq/http/client"
)

func Ticker(request Request) {
	ticker := time.NewTicker(time.Second * time.Duration(request.Duration))

	select {
	case <-ticker.C:
		switch request.Method {
		case "Post":
			client.Post(request.Url, request.Data, "application/json")
			break
		case "Get":
			url := request.Url
			if request.Data != nil {
				body, _ := json.Marshal(request.Data)
				url = url + string(body)
			}
			client.Get(url)
			break
		default:
			log.Println("wrong http method")
			break
		}
	}
	ticker.Stop()
}
