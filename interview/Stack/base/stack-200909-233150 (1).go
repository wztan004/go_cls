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
	p.size++
	return nil
}

func (p *stack) pop() (string, error) {
	var item string

	if p.top == nil {
		return "", errors.New("Empty Stack!")
	}

	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
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
func (p *stack) getSize() int {
	return p.size
}

func (p *stack) isEmpty() bool {
	return p.size == 0
}

func main() {

	myStack := &stack{nil, 0}
	fmt.Println("Created stack")
	fmt.Println()

	fmt.Print("Pushing nodes to the stack...\n\n")
	myStack.push("Mary")
	myStack.push("Jane")
	myStack.push("Xander")
	myStack.push("Marc")
	fmt.Println()

	tempStack := &stack{nil, 0}
	fmt.Println("Printing all the items in the stack...")
	for myStack.isEmpty() == false {

		item, _ := myStack.pop()
		tempStack.push(item)
		fmt.Println(item)

	}

	for tempStack.isEmpty() == false {
		item, _ := tempStack.pop()
		myStack.push(item)

	}

}
