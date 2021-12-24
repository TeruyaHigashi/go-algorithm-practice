package linkedlist

type CircularlyLinkedList struct {
	head   *node
	length int
}

func NewCircularlyLinkedList(elems ...interface{}) *CircularlyLinkedList {
	cll := &CircularlyLinkedList{}
	for _, elem := range elems {
		cll.AddEnd(elem)
	}
	return cll
}

func (cll *CircularlyLinkedList) AddBegin(elem interface{}) {
	added := &node{
		value:    elem,
		next:     cll.head,
		previous: nil,
	}
	if cll.head != nil {
		if cll.head.next == nil {
			cll.head.next = added
			cll.head.previous = added
			added.previous = cll.head
		} else {
			cll.head.previous.next = added
			cll.head.previous = added
			added.previous = cll.head.previous
		}
	}
	cll.head = added
	cll.length++
}

func (cll *CircularlyLinkedList) AddEnd(elem interface{}) {
	if cll.head == nil {
		cll.AddBegin(elem)
		return
	}
	added := &node{
		value:    elem,
		next:     cll.head,
		previous: nil,
	}
	if cll.head.next == nil {
		cll.head.next = added
		cll.head.previous = added
		added.previous = cll.head
	} else {
		cll.head.previous.next = added
		cll.head.previous = added
		added.previous = cll.head.previous
	}
	cll.length++
}

func (cll *CircularlyLinkedList) RemoveBegin() interface{} {
	if cll.head == nil {
		return nil
	}
	removed := cll.head
	if cll.head.next == nil {
		cll.head = nil
	} else {
		cll.head.next.previous = cll.head.previous
		cll.head.previous.next = cll.head.next
		cll.head = cll.head.next
	}
	removed.next = nil
	removed.previous = nil
	cll.length--
	return removed.value
}

func (cll *CircularlyLinkedList) RemoveEnd() interface{} {
	if cll.head == nil {
		return nil
	} else if cll.head.next == nil {
		return cll.RemoveBegin()
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
