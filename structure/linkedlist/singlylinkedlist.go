package linkedlist

type SinglyLinkedList struct {
	first  *node
	length int
}

func NewSinglyLinkedList(elems ...interface{}) *SinglyLinkedList {
	sentinel := &node{}
	sll := &SinglyLinkedList{first: sentinel, length: 0}
	for _, elem := range elems {
		sll.AddEnd(elem)
	}
	return sll
}

func (sll *SinglyLinkedList) AddBegin(elem interface{}) {
	added := &node{value: elem, next: nil}
	if sll.first != nil {
		added.next = sll.first
	}
	sll.first = added
	sll.length++
}

func (sll *SinglyLinkedList) AddEnd(elem interface{}) {
	if sll.first == nil {
		sll.AddBegin(elem)
		return
	}
	added := &node{value: &elem, next: nil}
	iter := sll.first
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = added
	sll.length++
}

func (sll *SinglyLinkedList) RemoveBegin() interface{} {
	if sll.first == nil {
		return nil
	}
	removed := sll.first
	sll.first = removed.next
	removed.next = nil
	sll.length--
	return removed.value
}

func (sll *SinglyLinkedList) RemoveEnd() interface{} {
	if sll.first == nil {
		return nil
	} else if sll.first.next == nil {
		return sll.RemoveBegin()
	}
	iter := sll.first
	for iter.next.next != nil {
		iter = iter.next
	}
	removed := iter.next
	removed.next = nil
	iter.next = nil
	sll.length--
	return removed.value
}

func (sll *SinglyLinkedList) Length() int {
	return sll.length
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.length == 0
}
