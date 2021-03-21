// https://courses.goschool.sg/courses/take/go-advanced/pdfs/16489897-go-school-go-advanced-practical-solutions

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

func checkBalancedParenthesis(expr string) bool {
	s := &stack{nil, 0}
	var x string

	// Traversing the Expression
	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' || expr[i] == '[' || expr[i] == '{' {
			// Push the element in the stack
			s.push(string(expr[i]))
			continue
		}

		// IF current current character is not opening
		// bracket, then it must be closing. So stack
		// cannot be empty at this point.
		if s.isEmpty() == true {
			return false
		}

		switch expr[i] {
		case ')':

			// Store the top element in a
			x, _ = s.pop()
			if x == "{" || x == "[" {
				return false
			}

		case '}':

			// Store the top element in b
			x, _ = s.pop()
			if x == "(" || x == "[" {
				return false
			}

		case ']':

			// Store the top element in c
			x, _ = s.pop()
			if x == "(" || x == "{" {
				return false
			}

		}
	}

	// Check Empty Stack
	return (s.isEmpty())
}

func main() {
	expr := "{()[}]"

	if checkBalancedParenthesis(expr) == true {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not balanced")
	}

}
