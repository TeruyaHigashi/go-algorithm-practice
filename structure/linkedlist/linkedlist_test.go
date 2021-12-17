package linkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	sut := NewSinglyLinkedList()
	lengthTest(sut, 0, t)

	sut = NewSinglyLinkedList()
	isEmptyTest(sut, true, t)

	sut = NewSinglyLinkedList()
	addBeginTest(sut, "[Go,]", t)

	addEndTest(sut, "[Go,]", t)

	sut = NewSinglyLinkedList("Go", true, 100)
	removeBeginTest(sut, "Go", "[true,100,]", t)

	sut = NewSinglyLinkedList("Go", false, 200)
	removeEndTest(sut, 200, "[Go,false,]", t)
}

func TestDoublyLinkedList(t *testing.T) {

	sut := NewDoublyLinkedList()
	lengthTest(sut, 0, t)

	sut = NewDoublyLinkedList()
	isEmptyTest(sut, true, t)

	sut = NewDoublyLinkedList()
	addBeginTest(sut, "[Go,]", t)

	addEndTest(sut, "[Go,]", t)

	sut = NewDoublyLinkedList("Go", true, 100)
	removeBeginTest(sut, "Go", "[true,100,]", t)

	sut = NewDoublyLinkedList("Go", false, 200)
	removeEndTest(sut, 200, "[Go,false,]", t)
}

func TestCircularlyLinkedList(t *testing.T) {

	sut := NewCircularlyLinkedList()
	lengthTest(sut, 0, t)

	sut = NewCircularlyLinkedList()
	isEmptyTest(sut, true, t)

	sut = NewCircularlyLinkedList()
	addBeginTest(sut, "[Go,]", t)

	addEndTest(sut, "[Go,]", t)

	sut = NewCircularlyLinkedList("Go", true, 100)
	removeBeginTest(sut, "Go", "[true,100,]", t)

	sut = NewCircularlyLinkedList("Go", false, 200)
	removeEndTest(sut, 200, "[Go,false,]", t)
}

func lengthTest(sut LinkedList, expected int, t *testing.T) {
	if sut.Length() != expected {
		t.Errorf("Lengthが期待値と相違")
	}
}

func isEmptyTest(sut LinkedList, expected bool, t *testing.T) {
	if sut.IsEmpty() != expected {
		t.Errorf("IsEmptyが期待値と相違")
	}
}

func addBeginTest(sut LinkedList, expectedJoined string, t *testing.T) {

	expectedLength := sut.Length() + 1
	sut.AddBegin("Go")

	if sut.Length() != expectedLength {
		t.Errorf("AddBegin後のlengthが期待値と相違")
	}
	if join(sut, ",") != expectedJoined {
		t.Errorf("AddBegin後の要素の並び順が期待値と相違")
	}
}

func addEndTest(sut LinkedList, expectedJoined string, t *testing.T) {

	expectedLength := sut.Length() + 1
	sut.AddEnd("Go")

	if sut.Length() != expectedLength {
		t.Errorf("AddEnd後のlengthが期待値と相違")
	}
	if join(sut, ",") != expectedJoined {
		t.Errorf("AddEnd後の要素の並び順が期待値と相違")
	}
}

func removeBeginTest(sut LinkedList, expectedRemoved interface{}, expectedJoined string, t *testing.T) {

	expectedLength := sut.Length() - 1
	removed := sut.RemoveBegin()

	if sut.Length() != expectedLength {
		t.Errorf("RemoveBegin後のlengthが期待値と相違")
	}
	if removed != expectedRemoved {
		t.Errorf("Removeした要素が期待値と相違")
	}
	if join(sut, ",") != expectedJoined {
		t.Errorf("RmoveBegin後の要素の並び順が期待値と相違")
	}
}

func removeEndTest(sut LinkedList, expectedRemoved interface{}, expectedJoined string, t *testing.T) {

	expectedLength := sut.Length() - 1
	removed := sut.RemoveEnd()

	if sut.Length() != expectedLength {
		t.Errorf("RemoveEnd後のlengthが期待値と相違")
	}
	if removed != expectedRemoved {
		t.Errorf("Removeした要素が期待値と相違")
	}
	if join(sut, ",") != expectedJoined {
		t.Errorf("RmoveEnd後の要素の並び順が期待値と相違")
	}
}

func join(linkedList LinkedList, separate string) string {
	joined := "["
	count := linkedList.Length()
	for i := 0; i < count; i++ {
		joined += fmt.Sprint(linkedList.RemoveBegin()) + separate
	}
	return joined + "]"
}
