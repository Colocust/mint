package queue

import "mint/container/linkedList"

type Queue struct {
	*linkedList.LinkedList
}

func NewQueue() *Queue {
	l := linkedList.NewLinkedList()
	q := &Queue{l}
	return q
}

func (q *Queue) DeQueue() interface{} {
	value, err := q.Shift()
	if err != nil {
		return nil
	}
	return value.Value
}

func (q *Queue) EnQueue(v interface{}) {
	q.Push(v)
}

func (q *Queue) Front() interface{} {
	top := q.Top()
	if top == nil {
		return nil
	}
	return top.Value
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}
