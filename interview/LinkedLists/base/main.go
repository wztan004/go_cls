// // 2: Node struct
// // 2: LinkedList struct
// // 17: addNode()
// // 10: printAllNodes()
// // Update base package if you changed something here

// package main

// import (
// 	"errors"
// 	"fmt"
// )

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

func (p *LinkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if index <= p.size && index >= 0 {
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

func (p *LinkedList) removeAtPos(index int) (string, error) {
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

	fmt.Println("Demoing get...")
	item, error := myList.get(1)

	if error != nil {
		fmt.Println("Invalid Index")
	} else {
		fmt.Println(item)
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
