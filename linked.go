package linked

// ╻  ╻┏┓╻╻┏ ┏━╸╺┳┓   ╻  ╻┏━┓╺┳╸┏━┓
// ┃  ┃┃┗┫┣┻┓┣╸  ┃┃   ┃  ┃┗━┓ ┃ ┗━┓
// ┗━╸╹╹ ╹╹ ╹┗━╸╺┻┛   ┗━╸╹┗━┛ ╹ ┗━┛

// A singly-linked list
type Node struct {
	value interface{}
	next  *Node
}

func EmptyList() Node {
	n := Node {
		value: nil,
		next: nil,
	}

	return n
}

func NewList(values ...interface{}) Node {
	list := EmptyList()

	for _, value := range values {
		list.Push(value)
	}

	return list
}

func (n Node) ToSlice() []interface{} {
	slice := []interface{}{}
	for {
		slice = append(slice, n.value)
		if n.next == nil {
			break
		}
		n.Next()
	}
	return slice
}

func (n *Node) Reverse() {
	curr := n
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		// This conditional prevents a cycle from forming
		if prev == nil {
			prev = &Node{ n.value, n.next }
		} else {
			prev = curr
		}
		curr = next
	}

	*n = *prev
}

func (n *Node) PushNode(node Node) {
	if n.value == nil {
		n.value = node.value
		return
	}

	if n.next == nil {
		n.next = &node
	} else {
		n.next.PushNode(node)
	}
}

func (n *Node) Push(v interface{}) {
	if n.value == nil {
		n.value = v
	} else if n.next == nil {
		n.next = &Node{ v, nil }
	} else {
		n.next.Push(v)
	}
}

func (n *Node) Length() int {
	if n.value == nil {
		return 0
	}

	length := 1

	if n.next == nil {
		return length
	}

	return length + n.next.Length()
}

func (n *Node) Next() {
	*n = *n.next
}

func (n Node) Index(i int) interface{} {
	for current := 0; current != i; current++ {
		n.Next()
	}
	return n.value
}
