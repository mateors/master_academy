package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	//https://pkg.go.dev/golang.org/x/crypto/bcrypt#pkg-constants
	//https://auth0.com/blog/hashing-in-action-understanding-bcrypt/
	//https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
	//https://yourbasic.org/golang/three-dots-ellipsis/
	//cost
	//min=4
	//max=31
	//default=10
	bs, err := bcrypt.GenerateFromPassword([]byte("test123"), 20)
	fmt.Println(string(bs), err)
}
