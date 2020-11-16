package linkedList

type (
	LinkedList struct {
		head *Node
		tail *Node
		size int
	}
	Node struct {
		value interface{}
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
		err := NewError("wrong index")
		return err
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
	node.value = value

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
	node := &Node{value: value}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		node.prev, l.tail.next, l.tail = l.tail, node, node
	}
	l.size++
	return
}

func (l *LinkedList) Unshift(value interface{}) {
	node := &Node{value: value}
	if l.size == 0 {
		l.head, l.tail = node, node
	} else {
		node.next, l.head.prev, l.head = l.head, node, node
	}
	l.size++
	return
}

func (l *LinkedList) Pop() (*Node, error) {
	if l.size == 0 {
		err := NewError("empty linkedList")
		return nil, err
	}
	node := l.tail
	node.prev.next, l.tail, node.prev = nil, node.prev, nil

	l.size--
	return node, nil
}

func (l *LinkedList) Shift() (*Node, error) {
	if l.size == 0 {
		err := NewError("empty linkedList")
		return nil, err
	}
	node := l.head
	node.next.prev, l.head, node.next = nil, node.next, nil

	l.size--
	return node, nil
}

func (l *LinkedList) Remove(index int) (*Node, error) {
	if index < 0 || index > l.size {
		err := NewError("wrong index")
		return nil, err
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

	next := prev.next
	prev.next, next.next.prev, next.next, next.prev = next.next, prev, nil, nil

	l.size--
	return next, nil
}

func (l *LinkedList) foreach() []int {
	var results []int

	current := l.head
	for current != nil {
		results = append(results, current.value.(int))
		current = current.next
	}

	return results
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
