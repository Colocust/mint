package task

import (
	"mint/container/linkedList"
	"sync"
)

type (
	MessageQueue struct {
		*linkedList.LinkedList
	}
	Message struct {
		Url       string `json:"url"`
		Content   string `json:"content"`
		RetryTime int
	}
)

var (
	instance *MessageQueue
	once     sync.Once
)

func GetInstance() *MessageQueue {
	once.Do(func() {
		l := linkedList.NewLinkedList()
		instance = &MessageQueue{l}
	})
	return instance
}

func (mq *MessageQueue) Product(m *Message) {
	mq.Push(m)
}

func (mq *MessageQueue) Consume() *Message {
	node, err := mq.Shift()
	if err != nil {
		return nil
	}
	return node.Value.(*Message)
}
