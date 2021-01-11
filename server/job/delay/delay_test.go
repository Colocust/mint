package delay

import (
	"fmt"
	"mint/server/job"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestDelay(t *testing.T) {
	heap := GetInstance()
	wg.Add(1)
	go func() {
		Boot(heap)
		defer wg.Done()
	}()

	node := &job.Node{Url: "s", Content: "s", Method: "s"}
	for i := 0; i < 1000000; i++ {
		now := int(time.Now().Unix())
		heap.Push(now, node)
	}
	wg.Wait()
	fmt.Println(heap.Size())
}
