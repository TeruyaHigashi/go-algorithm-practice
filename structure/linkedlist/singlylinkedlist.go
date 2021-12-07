package linkedlist

type SinglyLinkedList struct {
	head     *node
	sentinel *node
	length   int
}

func NewSinglyLinkedList(elems ...interface{}) *SinglyLinkedList {
	sentinel := &node{}
	sll := &SinglyLinkedList{head: sentinel, sentinel: sentinel, length: 0}
	for _, elem := range elems {
		sll.AddLast(elem)
	}
	return sll
}

func (sll *SinglyLinkedList) AddFirst(elem interface{}) {
	added := &node{value: &elem, next: nil}
	sll.length++
	if sll.head != sll.sentinel {
		added.next = sll.head
	} else {
		added.next = sll.sentinel
	}
	sll.head = added
}

func (sll *SinglyLinkedList) AddLast(elem interface{}) {
	added := &node{value: &elem, next: sll.sentinel}
	sll.length++
	if sll.head == sll.sentinel {
		sll.head = added
		return
	}
	iter := sll.head
	for iter.next != sll.sentinel {
		iter = iter.next
	}
	iter.next = added
}

func (sll *SinglyLinkedList) RemoveFirst() interface{} {
	if sll.Length() == 0 {
		return nil
	}
	sll.length--
	removed := sll.head
	sll.head = removed.next
	removed.next = nil
	return removed.value
}

func (sll *SinglyLinkedList) RemoveLast() interface{} {
	if sll.Length() == 0 {
		return nil
	} else if sll.Length() == 1 {
		return sll.RemoveFirst()
	}
	sll.length--
	iter := sll.head
	for iter.next.next != sll.sentinel {
		iter = iter.next
	}
	removed := iter.next
	removed.next = nil
	iter.next = sll.sentinel
	return removed.value
}

func (sll *SinglyLinkedList) Length() int {
	return sll.length
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.length == 0
}
