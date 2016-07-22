package data_structures

// Linked List Implementation

type Node struct{
	data int;
	next, prev *Node;

}

type List struct {
	head, tail *Node
	size int
}

/**
Implementation:
For Example: objectOfList.Add(5)
Returns: List after adding the int.
 */
func (l *List) Add(elem int) *List {
	n := &Node{data: elem}
	if l.head == nil {
		l.head = n		// First node
		l.size = 1;
	} else {
		l.tail.next = n	// Add after prev last node
		n.prev = l.tail // Link back to prev last node
		l.size++
	}
	l.tail = n
	return l
}

/**
Implementation:
For Example: objectOfList.Search(5)
Returns: List if found or Nil
 */
func (l *List) Search(elem int) *Node {

	found := false
	var ret *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.data == elem {
			found = true
			ret = n
		}
	}
	return ret
}

func (l *List) First() *Node {
	return l.head
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}
