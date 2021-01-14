package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap1 := new(Heap)
	for i := 10; i > 0; i-- {
		node := &Node{i, i}
		heap1.Add(node)
	}

	for k, v := range heap1.data {
		fmt.Println(k, v)
	}

	top, _ := heap1.Remove()
	fmt.Println(top)
	for k, v := range heap1.data {
		fmt.Println(k, v)
	}
}
