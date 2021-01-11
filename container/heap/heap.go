package heap

import (
	"errors"
	"strconv"
)

//最小堆，如果后续需要最大堆，这里会改结构体名
type (
	Heap struct {
		data []*Node
	}

	Node struct {
		Key   int
		Value interface{}
	}
)

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) IsEmpty() bool {
	return h.Size() == 0
}

//返回堆顶元素
func (h *Heap) Top() *Node {
	if h.IsEmpty() {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Add(n *Node) error {
	h.data = append(h.data, n)
	if err := h.siftUp(h.Size() - 1); err != nil {
		return err
	}

	return nil
}

func (h *Heap) Remove() (top *Node, err error) {
	top = h.Top()
	if top == nil {
		return
	}
	if err = h.swap(0, h.Size()-1); err != nil {
		return
	}

	h.data = h.data[0 : h.Size()-1]
	if err = h.siftDown(0); err != nil {
		return
	}
	return
}

func (h *Heap) siftDown(index int) error {
	for h.leftChild(index) < h.Size() {
		i := h.leftChild(index)
		if i+1 < h.Size() && h.data[i+1].Key < h.data[i].Key {
			i++
		}

		if h.data[i].Key >= h.data[index].Key {
			break
		}

		if err := h.swap(i, index); err != nil {
			return err
		}
		index = i
	}
	return nil
}

func (h *Heap) siftUp(index int) (err error) {
	for index > 0 {
		parent, err := h.parent(index)
		if err != nil {
			return err
		}
		if h.data[index].Key >= h.data[parent].Key {
			break
		}
		if err = h.swap(index, parent); err != nil {
			return err
		}
		index = parent
	}
	return nil
}

func (h *Heap) leftChild(index int) int {
	return index*2 + 1
}

func (h *Heap) rightChild(index int) int {
	return index*2 + 2
}

func (h *Heap) parent(index int) (parent int, err error) {
	if index == 0 {
		index := strconv.Itoa(index)
		err = errors.New(index + "don`t have parent")
		return
	}
	parent = (index - 1) / 2
	return
}

func (h *Heap) swap(i, j int) error {
	if i < 0 || j < 0 || i >= h.Size() || j >= h.Size() {
		return errors.New("wrong index")
	}
	h.data[i], h.data[j] = h.data[j], h.data[i]
	return nil
}

func NewHeap() *Heap {
	heap := new(Heap)
	return heap
}
