package main

import (
	"assignment4_cp3/datastruct"
	// "assignment4_cp3/utils"
	// "encoding/csv"
	// "encoding/hex"
	// "errors"
	"fmt"
	
	// "log"
	// "net/http"
	// "os"
	"time"
)


func main() {
	ll := datastruct.NewLinkedList()

	s1 := datastruct.Session{"ID1", "node1", time.Now()}
	ll.EnqueueSession(s1)
	
	s2 := datastruct.Session{"ID2", "node2", time.Now()}
	ll.EnqueueSession(s2)

	s3 := datastruct.Session{"ID3", "node3", time.Now()}
	ll.EnqueueSession(s3)


	ll.Remove("node3")

	s4 := datastruct.Session{"ID4", "node4", time.Now()}
	ll.EnqueueSession(s4)

	s, _ := ll.GetAllID()
	fmt.Println("All nodes", s)
	fmt.Println("size", ll.Size)
	fmt.Println("head", ll.Head.Session.Username)
	fmt.Println("tail", ll.Tail.Session.Username)
}


// Progress

// AuthenticateUser (ON HOLD, TO RESOLVE A SUB FUNCTION)
// Create New Session (DONE)
// Putting New Session To Linked List (DONE)
// COMPARE TIME (DONE)
// ADDING USER TO CURRENT SESSION (IN PROGRESS)
