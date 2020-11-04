package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
)

//from https://github.com/satori/go.uuid
func CreateUUID() (id string) {
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	return fmt.Sprintf("%s", u2)
}