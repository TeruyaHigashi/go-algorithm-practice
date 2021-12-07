package linkedlist

type DoublyLinkedList struct {
	head   *node
	length int
}

func (dll *DoublyLinkedList) AddFirst(elem interface{}) {
	added := &node{value: &elem, next: nil}
	dll.length++
	if dll.head != nil {
		dll.head.previous = added
		added.next = dll.head
	}
	dll.head = added
}
