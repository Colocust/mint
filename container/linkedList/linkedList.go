package linkedList

import "errors"

type (
	LinkedList struct {
		head *Node
		tail *Node
		size int
	}
	Node struct {
		Value interface{}
		prev  *Node
		next  *Node
	}
)

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	l.size = 0
	return l
}

func (l *LinkedList) Add(index int, value interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("wrong index")
	}
	if index == 0 {
		l.Unshift(value)
		return nil
	}
	if index == l.size {
		l.Push(value)
		return nil
	}

	node := new(Node)
	node.Value = value

	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	next := prev.next
	prev.next, node.prev, node.next, next.prev = node, prev, next, node

	l.size++
	return nil
}

func (l *LinkedList) Push(value interface{}) {
	node := &Node{Value: value}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		node.prev, l.tail.next, l.tail = l.tail, node, node
	}
	l.size++
	return
}

func (l *LinkedList) Unshift(value interface{}) {
	node := &Node{Value: value}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		node.next, l.head.prev, l.head = l.head, node, node
	}
	l.size++
	return
}

func (l *LinkedList) Pop() (node *Node, err error) {
	if l.size == 0 {
		err = errors.New("empty linkedList")
		return
	}
	node = l.tail
	prev := node.prev
	if prev == nil {
		l.head, l.tail = nil, nil
	} else {
		node.prev.next, l.tail, node.prev = nil, node.prev, nil
	}

	l.size--
	return node, nil
}

func (l *LinkedList) Shift() (node *Node, err error) {
	if l.size == 0 {
		err = errors.New("empty linkedList")
		return
	}
	node = l.head
	next := node.next
	if next == nil {
		l.head, l.tail = nil, nil
	} else {
		next.prev, l.head, node.next = nil, node.next, nil
	}

	l.size--
	return
}

func (l *LinkedList) Remove(index int) (next *Node, err error) {
	if index < 0 || index > l.size {
		err = errors.New("wrong index")
		return
	}

	if index == 0 {
		return l.Shift()
	}
	if index == l.size {
		return l.Pop()
	}

	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	next = prev.next
	prev.next, next.next.prev, next.next, next.prev = next.next, prev, nil, nil

	l.size--
	return
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Top() *Node {
	if l.size == 0 {
		return nil
	}
	return l.head
}

func (l *LinkedList) Bottom() *Node {
	if l.size == 0 {
		return nil
	}
	return l.tail
}
