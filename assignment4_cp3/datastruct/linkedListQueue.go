package datastruct

import (
	"fmt"
	"time"
	"log"
	"errors"
	"assignment4_cp3/constants"
)

// A SessionLinkedList is a linked list of user sessions, ordered by last
// logged in time, with the most recent session being at the Tail.
type SessionLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// A Node is used to describe a user's session details with a pointer to the 
// next user. The next user will always be someone who visited the website
// more recently. 
type Node struct {
	Session Session
	Next *Node
}

// NewSessionLinkedList returns an empty linked list for session management.
func NewSessionLinkedList() *SessionLinkedList {
    return &SessionLinkedList{nil,nil,0}
}

// EnqueueSession adds the user's latest session info to the linked list, but
// does not remove the previous session. Usually used with RemoveSession.
func (p *SessionLinkedList) EnqueueSession(s Session) {
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

// RemoveSession removes the user's session info from the linked list, usually
// used with EnqueueSession.
func (p *SessionLinkedList) RemoveSession(username string) {
	if p.Head == nil {
		fmt.Println("!3")
		return
	}

	if p.Size == 1 {
		if p.Head.Session.Username == username {
			p.Head = nil
			p.Tail = nil
			p.Size--
		}
		return
	}

	dummy := &Node{Session{"","",time.Now()}, p.Head}
	d := dummy

	// if matched in first node after dummy, i.e linked list's head 
	if d.Next.Session.Username == username {
		p.Head = d.Next.Next
		d.Next = d.Next.Next
		p.Size--
		fmt.Println("!1")
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
			fmt.Println("!2")
			return
		}
        d = d.Next
	}
	
}

// GetAllID (two L's one I) returns a list of all current active session IDs.
func (p *SessionLinkedList) GetAllID() ([]string, error) {
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

// CheckSessionID checks if the user's session ID is present in the linked list.
func (p *SessionLinkedList) CheckSessionID(id string) (bool, string) {
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

// hasSessionTimeout returns true if session Node has exceeded the timeout
// limit.
func hasSessionTimeout(t time.Time) bool {
	thirtyMinutes, err := time.ParseDuration(constants.Timeout)
	if err != nil {
		log.Fatalln("Error converting to Time format")
	}
	tNow := time.Now()
	diff := tNow.Sub(t)
	return diff > thirtyMinutes
}
