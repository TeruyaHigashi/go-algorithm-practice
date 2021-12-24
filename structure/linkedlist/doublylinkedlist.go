package linkedlist

type DoublyLinkedList struct {
	first  *node
	last   *node
	length int
}

func NewDoublyLinkedList(elems ...interface{}) *DoublyLinkedList {
	dll := &DoublyLinkedList{}
	for _, elem := range elems {
		dll.AddEnd(elem)
	}
	return dll
}

func (dll *DoublyLinkedList) AddBegin(elem interface{}) {
	added := &node{
		value:    elem,
		next:     nil,
		previous: nil,
	}
	if dll.first == nil {
		dll.last = added
	} else {
		dll.first.previous = added
		added.next = dll.first
	}
	dll.first = added
	dll.length++
}

func (dll *DoublyLinkedList) AddEnd(elem interface{}) {
	if dll.first == nil {
		dll.AddBegin(elem)
		return
	}
	added := &node{
		value:    elem,
		next:     nil,
		previous: dll.last,
	}
	dll.last.next = added
	dll.last = added
	dll.length++
}

func (dll *DoublyLinkedList) RemoveBegin() interface{} {
	if dll.first == nil {
		return nil
	}
	removed := dll.first
	if dll.first.next == nil {
		dll.first = nil
		dll.last = nil
	} else {
		dll.first = dll.first.next
		dll.first.previous = nil
	}
	removed.next = nil
	removed.previous = nil
	dll.length--
	return removed.value
}

func (dll *DoublyLinkedList) RemoveEnd() interface{} {
	if dll.first == nil {
		return nil
	} else if dll.first.next == nil {
		return dll.RemoveBegin()
	}
	removed := dll.last
	dll.last = dll.last.previous
	dll.last.next = nil
	removed.next = nil
	removed.previous = nil
	dll.length--
	return removed.value
}

func (dll *DoublyLinkedList) Length() int {
	return dll.length
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Length() == 0
}
