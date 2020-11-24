// 2: Node struct
// 3: Queue struct

// Mistakes
// 

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

type Queue struct {
	head *Node
	tail *Node
	size int 
}

func (q *Queue) enqueue(name string) {
	newNode := &Node{
		name,
		nil,
	}

	if q.head == nil {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}

	q.tail = newNode
	q.size++
}

func (q *Queue) printAllNodes() error {
	if q == nil {
		return errors.New("Nothing to print")
	}
	currentNode := q.head
	for currentNode.next != nil {
		fmt.Println(currentNode.item)
		currentNode = currentNode.next
	}
	fmt.Println(currentNode.item)
	return nil
}

func (q *Queue) dequeue() (string, error) {
	if q.head == nil {
		return "", errors.New("Empty Queue!")
	}

	item := q.head.item
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.size--
	return item, nil
}

func main() {
	myQueue := &Queue{nil, nil, 0}
	fmt.Println("Created queue")
	fmt.Println()
	fmt.Print("Enqueueing items to the queue...\n\n")
	myQueue.enqueue("Mary")
	myQueue.enqueue("Jane")
	myQueue.enqueue("Xander")
	myQueue.enqueue("Marc")
	fmt.Println("Showing all nodes in the queue...")
	myQueue.printAllNodes()
	fmt.Printf("There are %+v elements in the queue in totoal.\n", myQueue.size)

	myQueue.dequeue()
	fmt.Println("Showing all nodes in the queue...")
	myQueue.printAllNodes()
	fmt.Printf("There are %+v elements in the queue in totoal.\n", myQueue.size)
}