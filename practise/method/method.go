package main

import "fmt"

type Number int

func (a Number) duble() int {

	x := int(a) * 2
	return x
}

func main() {

	y := Number(10)
	fmt.Println(y.duble())
}
