// 2: Node struct
// 2: Stack struct
// 14: push
// 15: printAllNodes
// 14: pop

// Mistakes
// Visualized .next to wrong direction

// Use these variable names: Queue = q; string = name
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

type stack struct {
	top  *Node
	size int
}

func (p *stack) push(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}

	fmt.Println("Pushed", name)

	p.size++
	return nil
}

func (p *stack) printAllNodes() error {
	currentNode := p.top

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

func (p *stack) pop() (string, error) {
	if p.top == nil {
		return "", errors.New("Empty Stack!")
	}

	item := p.top.item

	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}

	p.size--
	return item, nil
}


func main() {
	// Section 1
	myStack := &stack{nil, 0}

	fmt.Println("Pushing nodes...")
	myStack.push("Mary")
	myStack.push("Jane")
	myStack.push("Xander")
	myStack.push("Marc")
	fmt.Println()

	// Section 2
	fmt.Println("Popped a node...")
	myStack.pop()
	myStack.printAllNodes()
}
