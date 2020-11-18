package message

import "mint/container/queue"

type (
	MQ struct {
		*queue.Queue
	}
	Message struct {
		Url       string
		Content   string
		RetryTime int
	}
)

func NewQueue() *MQ {
	return new(MQ)
}

func (q *MQ) Product(m Message) {
	q.EnQueue(m)
}

func (q *MQ) Consume() Message {
	m := q.DeQueue()
	return m.(Message)
}
