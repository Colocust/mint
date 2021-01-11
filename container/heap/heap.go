package heap

type (
	MinHeap struct {
		data []*MinHeapNode
	}

	MinHeapNode struct {
		Key   int
		Value interface{}
	}
)

func (h *MinHeap) Push(n *MinHeapNode) {
	h.data = append(h.data, n)
}

func NewMinHeap() *MinHeap {
	heap := new(MinHeap)
	return heap
}
