package main

import (
	"fmt"
	"sync"
)

func main() {

	igen := intGenerator()
	fmt.Println(">>", igen())
	fmt.Println(">>", igen())
	fmt.Println(">>", igen())
	//igen = nil

	//n := Make(func() { return 23 }) // Or something more expensive…
	//fmt.Println(n())                // Calculates the 23
	//fmt.Println(n() + 42)           // Reuses the calculated value

	//fmt.Println(addOrMultiply(true, add(4), multiply(4))) // 8
	//fmt.Println(addOrMultiply(false, add(4), multiply(4))) // 16
	fmt.Println(addOrMultiply2(true, add2, multiply2, 4))

}

type genType func() int

func intGenerator() genType {
	x := 0
	return func() int {
		x += 2
		return x
	}
}

func Make(f func() int) genType {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			v = f()
			f = nil // so that f can now be GC'ed
		})
		return v
	}
}

// func LazyAdd(a, b genType) genType {
// 	return Make(func() { return a() + b() })
// }

func add(x int) int {
	fmt.Println("executing add") // this is printed since the functions are evaluated first
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply") // this is printed since the functions are evaluated first
	return x * x
}

func addOrMultiply(add bool, onAdd, onMultiply int) int {
	if add {
		return onAdd
	}
	return onMultiply
}

//-----------higher-order-functions to rewrite this into a lazily evaluated version
func add2(x int) int {
	fmt.Println("executing add")
	return x + x
}

func multiply2(x int) int {
	fmt.Println("executing multiply")
	return x * x
}

// This is now a higher-order-function hence evaluation of the functions are delayed in if-else
func addOrMultiply2(add bool, onAdd, onMultiply func(t int) int, t int) int {
	if add {
		return onAdd(t)
	}
	return onMultiply(t)
}
