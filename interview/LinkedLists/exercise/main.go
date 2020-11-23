// 2: Node struct
// 2: LinkedList struct
// 17: addNode()
// 10: printAllNodes()

// Mistakes
// First try: get() did not check for index edge cases

// Use these variable names: LinkedList = p; Node = newNode; string = name
// Update base package if you changed something here

package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (p *LinkedList) addNode(name string) error {
	newNode := &Node{
		item:	name,
		next:	nil,
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

func (p *LinkedList) printAllNodes() {
	if p.head == nil {
		fmt.Println("nil")
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			fmt.Println(currentNode.item)
			currentNode = currentNode.next
		}
		fmt.Println(currentNode.item)
	}
}

func (p *LinkedList) get(index int) (string, error) {
	if p.head == nil {
		return "", errors.New("Empty Linked list!")
	}

	if index >= 0 && index < p.size {
		currentNode := p.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.next
		}
		item := currentNode.item
		return item, nil

	}
	return "", errors.New("Invalid Index")
}

func main() {
	// Section 1
	myList := &LinkedList{nil, 0}
	fmt.Println("Created linked list")
	fmt.Println()

	fmt.Print("Adding nodes to the linked list...\n\n")
	myList.addNode("Mary")
	myList.addNode("Jaina")
	myList.addNode("Xander")
	myList.addNode("Marc")
	fmt.Println("Showing all nodes in the linked list...")
	myList.printAllNodes()
	fmt.Printf("There are %+v elements in the list in totoal.\n", myList.size)
	fmt.Println()

	// Section 2
	fmt.Println("Demoing get...")
	item, error := myList.get(3)

	if error != nil {
		fmt.Println("Invalid Index")
	} else {
		fmt.Println(item)
	}

}