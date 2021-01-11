package server

import (
	"encoding/json"
	"sync"
)

var (
	mutex sync.Mutex
)

//func Consume() {
//	instance := job.GetInstance()
//	for {
//		mutex.Lock()
//		m := instance.Consume()
//		mutex.Unlock()
//
//		if m == nil {
//			continue
//		}
//
//		builder := client.NewBuilder().SetContent(m.Content).SetUrl(m.Url)
//		_, err := client.NewSender(builder).Send(netHttp.MethodPost)
//		if err != nil {
//			if m.RetryTime < retryMaxTime {
//				m.RetryTime++
//				mutex.Lock()
//				job.GetInstance().Product(m)
//				mutex.Unlock()
//			}
//		}
//	}
//}

func Delay(content string, resp *Response) {
	req := new(DelayRequest)
	if err := json.Unmarshal([]byte(content), req); err != nil {
		resp.Code = StatusServerError
		return
	}
	mutex.Lock()

	mutex.Unlock()
}

//func Add(request []string) *Response {
//	content, m := request[0], new(job.Message)
//
//	err := json.Unmarshal([]byte(content), m)
//	if err != nil {
//		return NewResponse(StatusArgsError, "Wrong Args")
//	}
//
//	mutex.Lock()
//	job.GetInstance().Product(m)
//	mutex.Unlock()
//
//	return NewResponse(StatusSuccess, "Success")
//}
