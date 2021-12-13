package linkedlist

type CircularlyLinkedList struct {
	head   *node
	length int
}

func (cll *CircularlyLinkedList) AddBegin(elem interface{}) {
	added := &node{
		value:    elem,
		next:     cll.head,
		previous: cll.head.previous,
	}
	cll.head.previous.next = added
	cll.head.previous = added
	cll.head = added
	cll.length++
}

func (cll *CircularlyLinkedList) AddEnd(elem interface{}) {
	if cll.head == nil {
		cll.AddBegin(elem)
	}
	added := &node{
		value:    elem,
		next:     cll.head,
		previous: cll.head.previous,
	}
	cll.head.previous.next = added
	cll.head.previous = added
	cll.length++
}

func (cll *CircularlyLinkedList) RemoveBegin() interface{} {
	if cll.head == nil {
		return nil
	}
	removed := cll.head
	cll.head.next.previous = cll.head.previous
	cll.head.previous.next = cll.head.next
	cll.head = cll.head.next
	removed.next = nil
	removed.previous = nil
	cll.length--
	return removed.value
}

func (cll *CircularlyLinkedList) RemoveEnd() interface{} {
	if cll.head == nil {
		return nil
	}
	removed := cll.head.previous
	cll.head.previous = cll.head.previous.previous
	cll.head.previous.previous.next = cll.head
	removed.next = nil
	removed.previous = nil
	cll.length--
	return removed.value
}

func (cll *CircularlyLinkedList) Length() int {
	return cll.length
}

func (cll *CircularlyLinkedList) IsEmpty() bool {
	return cll.Length() == 0
}
