package datastruct

import (
	"fmt"
	"time"
	"log"
	"errors"

)

type Node struct {
	Session Session
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

//New
func NewLinkedList() *LinkedList {
    return &LinkedList{nil,nil,0}
}


// Does not include removing existing session by the user
func (p *LinkedList) EnqueueSession(s Session) {
	// includes timestamp of the new node
	mNode := &Node{s, nil}

	if p.Head == nil {
		p.Head = mNode
		p.Tail = mNode
	} else {
		p.Tail.Next = mNode
		p.Tail = mNode
		fmt.Println("Tail is now", p.Tail.Session.Username)
	}
	p.Size++
}

func (p *LinkedList) DeleteSession(s string) {
	currentNode := p.Head
	for currentNode.Next != nil {
        if currentNode.Next.Session.SessionUUID == s {
			currentNode.Next = currentNode.Next.Next
			p.Size--
        }
        currentNode = currentNode.Next
	}
}



func (p *LinkedList) Remove(username string) {
	if p.Head == nil {
		return
	}

	if p.Size == 1 {
		p.Head = nil
		p.Tail = nil
		p.Size--
		return
	}

	dummy := &Node{Session{"","",time.Now()}, p.Head}
	d := dummy

	// if matched in first node after dummy, i.e linked list's head 
	if d.Next.Session.Username == username {
		p.Head = d.Next.Next
		d.Next = d.Next.Next
		p.Size--
		return
	}

    for d.Next != nil {
        if d.Next.Session.Username == username {
			if d.Next.Next == nil {
				p.Tail = d
			}
			d.Next = d.Next.Next
			p.Size--
			fmt.Println(d.Session.Username)
			return
		}
        d = d.Next
	}
	
}


func (p *LinkedList) GetAllID() ([]string, error) {
	currentNode := p.Head

	if currentNode == nil {
		return nil, errors.New("No items found")
	}

	var output []string
	for currentNode.Next != nil {
		output = append(output, currentNode.Session.SessionUUID)
		currentNode = currentNode.Next
	}
	output = append(output, currentNode.Session.SessionUUID)

	return output, nil
}

// CheckSessionID checks if user's session ID is already active in the linked list
func (p *LinkedList) CheckSessionID(id string) (bool, string) {
	currentNode := p.Head

	if currentNode == nil {
		return false, ""
	}

	if (currentNode.Session.SessionUUID == id && !hasSessionTimeout(currentNode.Session.CreatedAt)) {
		return true, currentNode.Session.Username
	}

	for currentNode.Next != nil {
		if (currentNode.Session.SessionUUID == id && !hasSessionTimeout(currentNode.Session.CreatedAt)) {
			return true, currentNode.Session.Username
		}
		currentNode = currentNode.Next
	}
	return false, ""
}



func hasSessionTimeout(t time.Time) bool {
	thirtyMinutes, err := time.ParseDuration("30m")
	if err != nil {
		log.Fatalln("Error converting to Time format")
	}
	tNow := time.Now()
	diff := tNow.Sub(t)
	return diff > thirtyMinutes
}
