//https://courses.goschool.sg/courses/take/go-advanced/pdfs/16489897-go-school-go-advanced-practical-solutions

package main

import (
	"errors"
	"fmt"
)

type Node struct {
	priority int
	item     string
	next     *Node
}
type queue struct {
	front *Node
	back  *Node
	size  int
}

func (p *queue) enqueue(name string, prty int) error {
	newNode := &Node{
		priority: prty,
		item:     name,
		next:     nil,
	}
	if p.front == nil {
		p.front = newNode
	} else {
		if p.front.priority > prty {
			// Insert New Node before front
			newNode.next = p.front
			p.front = newNode
		} else {
			currentNode := p.front
			for currentNode.next != nil && currentNode.next.priority <= prty {
				currentNode = currentNode.next
			}
			// Either at the ends of the list
			// or at required position
			newNode.next = currentNode.next
			currentNode.next = newNode
		}
	}
	p.size++
	return nil
}
func (p *queue) dequeue() (string, error) {
	var item string
	if p.front == nil {
		return "", errors.New("empty queue!")
	}
	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}
func (p *queue) printAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}
func (p *queue) isEmpty() bool {
	return p.size == 0
}
func main() {
	myQueue := &queue{nil, nil, 0}
	fmt.Println("Created queue")
	fmt.Println()
	fmt.Print("Enqueueing items to the queue...\n\n")
	myQueue.enqueue("Mary", 3)
	myQueue.enqueue("John", 2)
	myQueue.enqueue("Jane", 1)
	myQueue.enqueue("Xander", 3)
	myQueue.enqueue("Marc", 2)
	fmt.Println("Showing all nodes in the queue...")
	myQueue.printAllNodes()
	fmt.Printf("There are %+v elements in the queue in totoal.\n", myQueue.size)
}
