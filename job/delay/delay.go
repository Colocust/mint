package job

import (
	"mint/container/heap"
	"sync"
)

type (
	DelayJob struct {
		Url     string `json:"url"`
		Content string `json:"content"`
		Method  string `json:"method"`
	}
	DelayJobQueue struct {
		*heap.Heap
	}
)

var (
	delay     *DelayJobQueue
	delayOnce sync.Once
)

func (queue *DelayJobQueue) Product(when int, job *DelayJob) {
	node := &heap.Node{Key: when, Value: job}
	queue.Add(node)
}

func GetDelayJob() *DelayJobQueue {
	delayOnce.Do(func() {
		h := heap.NewHeap()
		delay = &DelayJobQueue{h}
	})
	return delay
}
