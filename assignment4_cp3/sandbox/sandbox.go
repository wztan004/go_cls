// Used for testing purposes. Ignore this package/file.
package main

import (
	"fmt"
	"strings"
)

func main() {
	q := `confidential/venue_202009.csv`
	fmt.Println(strings.Contains(q, "confidential/venues_"))
}