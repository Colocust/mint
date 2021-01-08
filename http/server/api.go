package server

import (
	"encoding/json"
	"mint/http/client"
	"mint/job"
	netHttp "net/http"
	"sync"
)

var (
	mutex        sync.Mutex
	retryMaxTime int
)

//func init() {
//	retryMaxTime = int(config.Read("retry_max_time").(float64))
//}

func Add(request []string) *Response {
	content, m := request[0], new(job.Message)

	err := json.Unmarshal([]byte(content), m)
	if err != nil {
		return NewResponse(StatusArgsError, "Wrong Args")
	}

	mutex.Lock()
	job.GetInstance().Product(m)
	mutex.Unlock()

	return NewResponse(StatusSuccess, "Success")
}

func Consume() {
	instance := job.GetInstance()
	for {
		mutex.Lock()
		m := instance.Consume()
		mutex.Unlock()

		if m == nil {
			continue
		}

		builder := client.NewBuilder().SetContent(m.Content).SetUrl(m.Url)
		_, err := client.NewSender(builder).Send(netHttp.MethodPost)
		if err != nil {
			if m.RetryTime < retryMaxTime {
				m.RetryTime++
				mutex.Lock()
				job.GetInstance().Product(m)
				mutex.Unlock()
			}
		}
	}
}

func Delay(req []string, resp Response) {

}

func Fixed(req []string, resp Response) {

}

func Ticker(req []string, resp Response) {

}
