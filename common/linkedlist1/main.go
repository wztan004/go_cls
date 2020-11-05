package main

import (
	"fmt"
)


type Node struct {
    Next *Node
	Val  int
	Size int
}

func (n *Node) Append(val int) {
    end := &Node{Val: val}
    here := n
    for here.Next != nil {
        here = here.Next
    }
	here.Next = end
	n.Size++
}

func Remove(n *Node, val int) *Node {
	if n == nil {
		return n
	}
	var p1 Node
	p1.Next = n
	p2 := &p1
	for n != nil {
		if n.Val == val {
			p2.Next, n = n.Next, n.Next
		} else {
			p2, n = n, n.Next
		}
	}
	return p1.Next
}

func NewNode(val int) *Node {
    return &Node{Val: val}
}

func main() {
    n := NewNode(1)
    n.Append(2)
	n.Append(3)
    n.Append(4)
    n.Append(5)

    m := Remove(n, 1)

    for m != nil {
        fmt.Println(m.Val)
        m = m.Next
	}
	fmt.Println(n,n.Val)
}