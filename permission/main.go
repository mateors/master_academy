package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func init() {

}

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := user.Username
	fmt.Printf("Username: %s\n", username)

	evar1 := os.Getenv("MY_ENV1")
	evar2 := os.Getenv("SLOG")
	fmt.Println("evar1:", evar1)
	fmt.Println("evar2:", evar2)

	list := os.Environ()
	for i, evar := range list {
		fmt.Println("listEnviron:", i, evar)
	}

	createDirectory()
}

func createDirectory() {

	_ = os.MkdirAll("tdata/1.txt", 0757)
	//syscall.Umask(0)
	_ = os.MkdirAll("/tmp/dirs/2", 0664)
}
