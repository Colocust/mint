package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"tiny-tq/src/parameter"
	"tiny-tq/tiny/http/client"
)

func Ticker(request parameter.Request) {
	ticker := time.NewTicker(time.Second * time.Duration(request.Duration))

	select {
	case <-ticker.C:
		builder := new(client.NetBuilder)
		content, _ := json.Marshal(request.Data)

		switch request.Method {
		case "Post":
			resp, _ := builder.SetUrl(request.Url).SetContent(string(content)).NewNetSender().Send(http.MethodPost)
			fmt.Println(resp)
			break
		case "Get":
			resp, _ := builder.SetUrl(request.Url).SetContent(string(content)).NewNetSender().Send(http.MethodGet)
			fmt.Println(resp)
			break
		default:
			log.Println("wrong http method")
			break
		}
	}
	ticker.Stop()
}

