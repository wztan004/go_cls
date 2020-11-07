package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"assignment4_cp3/datastruct"
	"time"
	"golang.org/x/crypto/blake2b"
)

//from https://github.com/satori/go.uuid
func CreateUUID() (id string) {
	mUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	return fmt.Sprintf("%s", mUUID)
}


func CreateSessionStruct(username string) (datastruct.Session) {
	mUUID := CreateUUID()
	mSession := datastruct.Session{mUUID, username, time.Now()}
	return mSession
}



func CreateChecksum(plain string) (checksum string) {
	b := []byte(plain)
	h := blake2b.Sum256(b)
	return fmt.Sprintf("%x", h)
}
