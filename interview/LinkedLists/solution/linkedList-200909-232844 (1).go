// seems to be an issue with get() index and addAtPos index. It seems to be 1-based index instead of 0-based inedx
// I've edited official solution in get() and addAtPos
// If you're using this for development, thoroughly test these first! E.g. negative int as argument for certain methods.

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

func (p *linkedList) get(index int) (string, error) {
	if p.head == nil {
		return "", errors.New("Empty Linked list!")
	}
	if index < p.size {
		currentNode := p.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.next
		}
		item := currentNode.item
		return item, nil

	}
	return "", errors.New("Invalid Index")
}

func (p *linkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	// if equals, means append
	if index <= p.size {
		if index == 0 {
			newNode.next = p.head
			p.head = newNode

		} else {
			currentNode := p.head
			var prevNode *Node
			for i := 0; i < index; i++ {
				prevNode = currentNode
				currentNode = currentNode.next
			}
			newNode.next = currentNode
			prevNode.next = newNode

		}
		p.size++
		return nil
	} else {
		return errors.New("Invalid Index")
	}
}

func (p *linkedList) remove(index int) (string, error) {
	var item string

	if p.head == nil {
		return "", errors.New("Empty Linked list!")
	}
	if index >= 0 && index < p.size {
		if index == 1 {
			item = p.head.item
			p.head = p.head.next
		} else {
			var currentNode *Node = p.head
			var prevNode *Node
			for i := 0; i < index; i++ {
				prevNode = currentNode
				currentNode = currentNode.next

			}
			item = currentNode.item
			prevNode.next = currentNode.next
		}
	}
	p.size--
	return item, nil
}



func main() {

	myList := &linkedList{nil, 0}
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

	fmt.Println("Demoing get...")
	item, error := myList.get(1)

	if error == nil {

		fmt.Println(item)
	} else {
		fmt.Println("Invalid Index")
	}
	fmt.Println()
	fmt.Println("Adding at index...")

	myList.addAtPos(4, "Chris")
	myList.printAllNodes()


	fmt.Println()
	fmt.Println("Removing at index...")
	myList.remove(2)
	myList.printAllNodes()

}
