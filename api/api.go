package api

import (
	"encoding/json"
	"mint/config"
	"mint/http"
	"mint/task"
	http2 "net/http"
	"sync"
)

var (
	mutex        sync.Mutex
	retryMaxTime int
)

func init() {
	retryMaxTime = int(config.Read("retry_max_time").(float64))
}

func Add(request []string) *Response {
	content, m := request[0], new(task.Message)

	err := json.Unmarshal([]byte(content), m)
	if err != nil {
		return NewResponse(StatusArgsError, "Wrong Args")
	}

	mutex.Lock()
	task.GetInstance().Product(m)
	mutex.Unlock()

	return NewResponse(StatusSuccess, "Success")
}

func Consume() {
	instance := task.GetInstance()
	for {
		mutex.Lock()
		m := instance.Consume()
		mutex.Unlock()

		if m == nil {
			continue
		}

		builder := http.NewBuilder()
		builder.SetContent(m.Content).SetUrl(m.Url)

		sender := http.NewSender(builder)
		_, err := sender.Send(http2.MethodPost)
		if err != nil {
			if m.RetryTime < retryMaxTime {
				m.RetryTime++
				mutex.Lock()
				task.GetInstance().Product(m)
				mutex.Unlock()
			}
		}
	}
}
