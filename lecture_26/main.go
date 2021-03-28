package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	//password encrypt
	bs, err := bcrypt.GenerateFromPassword([]byte("test123"), 14)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(bs))

	//password compare
	err = bcrypt.CompareHashAndPassword(bs, []byte("test123"))
	fmt.Println(err)
}
