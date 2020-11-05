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

	s := datastruct.Session{"ID1", "node1", time.Now()}
	ll.EnqueueSession(s)

	s = datastruct.Session{"ID2", "node2", time.Now()}
	ll.EnqueueSession(s)

	s = datastruct.Session{"ID3", "node3", time.Now()}
	ll.EnqueueSession(s)

	ll.Remove("ID1")

	s1, _ := ll.GetAllID()
	fmt.Println(s1)

}


// Progress

// AuthenticateUser (ON HOLD, TO RESOLVE A SUB FUNCTION)
// Create New Session (DONE)
// Putting New Session To Linked List (DONE)
// COMPARE TIME (DONE)
// ADDING USER TO CURRENT SESSION (IN PROGRESS)
