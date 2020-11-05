package datastruct

import (
	"testing"
	"time"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()

	s := Session{"ID1", "node1", time.Now()}
	ll.EnqueueSession(s)

	e := ll.CheckSessionID("ID1")
	if e != true {
		t.Errorf("CheckSessionID: Expecting %v, got %v", !e, e)
	}

	s = Session{"ID2", "node2", time.Now()}
	ll.EnqueueSession(s)

	s = Session{"ID3", "node3", time.Now()}
	ll.EnqueueSession(s)

	sizeList, err := ll.GetAllID()
	if err != nil {
		t.Errorf("Sub-func: Encountered errors from sub-function")
	}

	e1, r1 := ll.Size, len(sizeList)
	if (e1 != r1) {
		t.Errorf("Len(GetAllID) == ll.Size: Expecting %v, got %v", e1, r1)
	}

	e2, r2 := "node1", ll.Head.Session.Username
	if (e2 != r2) {
		t.Errorf("Head: Expecting %v, got %v", e2, r2)
	}

	e3, r3 := "node3", ll.Tail.Session.Username
	if (e3 != r3) {
		t.Errorf("Tail: Expecting %v, got %v", e3, r3)
	}

	ll.Remove("ID1")

	e4, r4 := 2, ll.Size
	if (e4 != r4) {
		t.Errorf("DeleteSession: Expecting %v, got %v", e4, r4)
	}
}