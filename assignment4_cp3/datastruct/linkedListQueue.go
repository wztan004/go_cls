package datastruct

import (
	"time"
	"log"
	"errors"
	"fmt"
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
	// In addition, also adds node based on time
	mNode := &Node{s, nil}

	if p.Head == nil {
		p.Head = mNode
		p.Tail = mNode
	} else {
		p.Tail.Next = mNode
		p.Tail = mNode
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



func (p *LinkedList) Remove(s string) {
	dummy := &Node{Session{"","",time.Now()}, p.Head}
	d := dummy

	// if matched in first node after dummy, i.e linked list's head 
	if d.Next.Session.SessionUUID == s {
		p.Head = d.Next.Next
		d.Next = d.Next.Next
		p.Size--
		return
	}

    for d.Next != nil {
        if d.Next.Session.SessionUUID == s {
			fmt.Println(d.Next.Session.SessionUUID)
			fmt.Println(s)
			d.Next = d.Next.Next
			p.Size--
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
func (p *LinkedList) CheckSessionID(id string) (bool) {
	currentNode := p.Head

	if currentNode == nil {
		return false
	}

	if (currentNode.Session.SessionUUID == id) {
		return true
	}

	for currentNode.Next != nil {
		if (currentNode.Session.SessionUUID == id) {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
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

// func (p *LinkedList) Remove(i int) {
// 	currentNode := p.Head
// 	j := 0

// 	for j < i-1 {
// 		j++
// 		// Stop at the node before the to-be-removed node
// 		currentNode = currentNode.Next
// 	}

// 	remove := currentNode.Next
// 	currentNode.Next = remove.Next

// 	p.Size--
// }



// Commented out as adding node only occurs at the back
// func (p *LinkedList) AddNode(mNodeq Node) {
// 	// In addition, also adds node based on time
// 	mNode := &mNodeq

// 	if p.Head == nil {
// 		p.Head = mNode
// 	} else {
// 		currentNode := p.Head
// 		if mNode.Session.CreatedAt <= currentNode.Session.CreatedAt {
// 			mNode.Next = currentNode
// 			p.Head = mNode
// 		} else {
// 			for currentNode.Next != nil && mNode.Date >= currentNode.Next.Date {
// 				currentNode = currentNode.Next
// 			}
// 			mNode.Next = currentNode.Next
// 			currentNode.Next = mNode
// 		}
// 	}
// 	p.Size++
// }


// Commented out as this func is not necessary
// func (p *LinkedList) Get(i int) *Node {
// 	currentNode := p.Head
// 	j := 0

// 	for j < i {
// 		j++
// 		currentNode = currentNode.Next
// 	}

// 	return currentNode
// }

