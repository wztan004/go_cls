// Implement addNode, printAllNodes, remove(), addAtPos()
// Refresher on linked list construction
// Mistakes: missed out a section in printAllNodes
// Go Advanced

// Solution seems to be an issue with get() index return val, so I removed this function from self-test

package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}





func main() {
	h := &Node{"head", nil}
	ll := linkedList{h, 1}
	ll.addNode("1")
	ll.addNode("2")
	ll.addNode("3")
	ll.printAllNodes()
}