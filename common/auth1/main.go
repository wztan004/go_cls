// create a signed token
// https://godoc.org/github.com/dgrijalva/jwt-go

package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("f")) //our secret
	fmt.Println(signedToken)
}