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
		next:     dll.first,
		previous: nil,
	}
	dll.first.previous = added
	dll.first = added
	if dll.last == nil {
		dll.last = added
	}
	dll.length++
}

func (dll *DoublyLinkedList) AddEnd(elem interface{}) {
	if dll.first == nil {
		dll.AddBegin(elem)
		return
	}
	added := &node{
		value:    &elem,
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
	dll.first = dll.first.next
	dll.first.previous = nil
	removed.next = nil
	removed.previous = nil
	dll.length--
	return removed.value
}

func (dll *DoublyLinkedList) RemoveEnd() interface{} {
	if dll.first == nil {
		return nil
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
