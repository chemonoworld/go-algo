package c4_linked_list

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) pushBack(value T) {
	curNode := l.root
	for {
		if curNode.next != l.tail {
			curNode = curNode.next
		} else {
			break
		}
	}

	curNode.next = &Node[T]{
		l.tail,
		value,
	}
}
