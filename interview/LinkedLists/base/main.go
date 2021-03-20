// 14: get()
// 25: addAtPos()
// 28: removeAtPos()

// Mistakes
// First try: get() did not check for index edge cases

// Use these variable names: LinkedList = p; Node = newNode; string = name
// Update base package if you changed something here

package main

import (
	"fmt"
	"errors"
)

type Node struct {
	item string
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) addNode (item string) {
	newNode := &Node{
		item,
		nil,
	}

	if list.head == nil {
		list.head = newNode
	} else {
		currNode := list.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
	}

	list.size++
}

func (list *LinkedList) printAllNodes () {
	if list.head == nil {
		fmt.Println("Blank!")
	} else {
		currNode := list.head
		for currNode.next != nil {
			fmt.Println(currNode.item)
			currNode = currNode.next
		}
		fmt.Println(currNode.item)
	}
}

func (list *LinkedList) get(index int) (string, error) {
	if index < 0 || index >= list.size {
		return "", errors.New("Invalid index")
	}

	currNode := list.head
	for i:=0; i < list.size; i++ {
		if i == index {
			return currNode.item, nil
		}
		currNode = currNode.next
	}
	return "", errors.New("Not found")
}

func (list *LinkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if index <= list.size && index >= 0 {
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			currentNode := list.head
			var prevNode *Node
			for i := 0; i < index; i++ {
				prevNode = currentNode
				currentNode = currentNode.next
			}
			newNode.next = currentNode
			prevNode.next = newNode
		}

		list.size++
		return nil
	} else {
		return errors.New("Invalid Index")
	}
}

func (list *LinkedList) removeAtPos(index int) (string, error) {
	if index < 0 || index >= list.size {
		return "", errors.New("Invalid index")
	}
	if index == 0 {
		removedNode := list.head
		list.head = list.head.next
		return removedNode.item, nil
	}

	currNode := list.head
	for currNode.next != nil {
		
	}

	list.size--
	return "", errors.New("Invalid")
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
	item, error := myList.get(4)

	if error != nil {
		fmt.Println("Invalid Index")
	} else {
		fmt.Println(item)
	}

	// Section 3
	fmt.Println()
	fmt.Println("Adding at index...")
	err := myList.addAtPos(0, "Chris")

	if err != nil {
		fmt.Println(err)
	}

	myList.printAllNodes()


	// // Section 4
	// fmt.Println()
	// fmt.Println("Removing at index...")
	// x, err := myList.removeAtPos(5)
	
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Removed", x)
	// }

	// myList.printAllNodes()
}