package datastruct

import (
	"testing"
	"time"
)

// Tests for Session{}, NewLinkedList(), EnqueueSession(), CheckSessionID()
func TestLinkedList_1(t *testing.T) {
	ll := NewLinkedList()

	s := Session{"ID1", "node1", time.Now()}
	ll.EnqueueSession(s)

	_, e := ll.CheckSessionID("ID1")
	r := "node1"
	if e != r {
		t.Errorf("CheckSessionID: Expecting %v, got %v", e, r)
	}
}

// Tests for Session.Head, Session.Tail, GetAllID(), Remove()
func TestLinkedList_2(t *testing.T) {
	// Start: Boilerplate
	ll := NewLinkedList()
	s1 := Session{"ID1", "node1", time.Now()}
	ll.EnqueueSession(s1)
	s2 := Session{"ID2", "node2", time.Now()}
	ll.EnqueueSession(s2)
	s3 := Session{"ID3", "node3", time.Now()}
	ll.EnqueueSession(s3)
	// End: Boilerplate

	sizeList, err := ll.GetAllID()
	if err != nil {
		t.Errorf("Sub-func: Encountered errors from sub-function")
	}

	e1, r1 := ll.Size, len(sizeList)
	if (e1 != r1) {
		t.Errorf("Len(GetAllID) == ll.Size: Expecting %v, got %v", e1, r1)
	}


	ll.Remove("node1")

	e2, r2 := "node2", ll.Head.Session.Username
	if (e2 != r2) {
		t.Errorf("Head: Expecting %v, got %v", e2, r2)
	}

	ll.Remove("node3")

	e3, r3 := "node2", ll.Tail.Session.Username
	if (e3 != r3) {
		t.Errorf("Tail: Expecting %v, got %v", e3, r3)
	}

	e4, r4 := 1, ll.Size
	if (e4 != r4) {
		t.Errorf("Remove: Expecting %v, got %v", e4, r4)
	}
}