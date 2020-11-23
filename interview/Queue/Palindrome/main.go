// https://courses.goschool.sg/courses/take/go-advanced/pdfs/16489897-go-school-go-advanced-practical-solutions

package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	var q *queue
	var s *stack
	isPalindrome := true
	text = "level"
	text = strings.ToLower(strings.Trim(text, " "))
	q = &queue{nil, nil, 0}
	s = &stack{nil, 0}

	for i := 0; i < len(text); i++ {
		ch := string(text[i])
		q.enqueue(ch)
		s.push(ch)
	}

	for i := 0; i < len(text); i++ {
		ch1, _ := q.dequeue()
		ch2, _ := s.pop()
		if ch1 != ch2 {
			isPalindrome = false
			break
		}
	}
	
	if isPalindrome {
		fmt.Println("A palindrome!")
	} else {
		fmt.Println("Not a palindrome")
	}
}
