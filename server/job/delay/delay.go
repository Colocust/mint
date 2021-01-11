package delay

import (
	"fmt"
	"mint/server/container/heap"
	"mint/server/job"
	"sync"
	"time"
)

type (
	Queue struct {
		*heap.Heap
	}
)

var (
	queue *Queue
	once  sync.Once
	mutex sync.Mutex
)

func (queue *Queue) Push(when int, n *job.Node) error {
	node := &heap.Node{Key: when, Value: n}

	mutex.Lock()
	err := queue.Add(node)
	mutex.Unlock()

	if err != nil {
		return err
	}
	return nil
}

func (queue *Queue) scan() {
	i := 0
	for {
		mutex.Lock()
		top := queue.Top()
		mutex.Unlock()

		if top != nil && int64(top.Key) <= time.Now().Unix() {
			mutex.Lock()
			top, _ = queue.Remove()
			mutex.Unlock()

			i++
			fmt.Println(i)
		}
	}
}

func GetInstance() *Queue {
	once.Do(func() {
		h := heap.NewHeap()
		queue = &Queue{h}
	})
	return queue
}

func Boot(queue *Queue) {
	queue.scan()
}
