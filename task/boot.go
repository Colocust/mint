package task

import (
	"encoding/json"
	"mint/config"
	"mint/http/client"
)

var ch chan *Task

func Boot() {
	taskNum := config.Read("task_num").(float64)
	ch = make(chan *Task, int(taskNum))
	consumer()
}

func consumer() {
	for {
		t := <-ch
		content, _ := json.Marshal(t.data)

		builder := new(client.NetBuilder)
		sender := builder.SetUrl(t.url).SetContent(string(content)).NewNetSender()
		sender.Send("POST")
	}
}

func product(t *Task) {
	ch <- t
}

type (
	Task struct {
		url  string
		data map[string]interface{}
	}
)
