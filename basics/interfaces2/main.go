// https://www.alexedwards.net/blog/interfaces-explained

package main

import (
    "fmt"
    "strconv"
    "log"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    log.Println(s.String(), "writeLog")
}

func main() {
    // Initialize a Count object and pass it to WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(3)
    WriteLog(count)
}