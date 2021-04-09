package main

import (
	"fmt"
	"runtime"
)

func main() {
	//set GOARCH=386
	//windows = 386
	//compilation for win32 bit operating system
	fmt.Printf("%s = %s\n", runtime.GOOS, runtime.GOARCH)
}
