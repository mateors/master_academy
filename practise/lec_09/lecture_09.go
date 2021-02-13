package main

import (
	"fmt"
)

func main() {

	//primitive data types == আদিম
	//Built-in Types
	//Basic types
	//স্বরবর্ণ
	//https://bn.wikipedia.org/wiki/%E0%A6%AC%E0%A6%BE%E0%A6%82%E0%A6%B2%E0%A6%BE_%E0%A6%B8%E0%A7%8D%E0%A6%AC%E0%A6%B0%E0%A6%AC%E0%A6%B0%E0%A7%8D%E0%A6%A3

	/*
	   In computer science, a primitive is a fundamental data type that cannot be broken down
	   into a more simple data type. For example, an integer is a primitive data type, while an array,
	   which can store multiple data types, is not.
	*/

	var char rune
	char = 'A'
	//char = 66
	fmt.Printf("%c\n", char)
	fmt.Printf("%T", char)

	//alias
	type CHARARRAY = string
	var ca CHARARRAY
	fmt.Printf("%T %v", ca, ca)

	//mini programs
	// var name string
	// var age int
	// fmt.Print("Enter your name: ")
	// fmt.Scanf(`%s`, &name) //variable name's memory address
	// //fmt.Scanf(`%s %d`, &name, &age)
	// fmt.Printf("%s is %d years old\n", name, age)
	// fmt.Println(&age)

	//fmt.Println(char)
	//numbers := []int{10, 20, 30, 40, 50}
	//numbers := make([]int, 5, 10)
	//fmt.Println(numbers, cap(numbers))
	//In computer science, a literal is a notation for representing
	//a fixed value in source code

	//fmt.Printf("%T\n", numbers)
	//fmt.Println(reflect.TypeOf(numbers).Kind().String())

	//var numbers [3]int
	//numbers := [3]int{}
	numbers := make([]int, 0)

	fmt.Println(numbers)
	// numbers := make([]int, 0)
	// for i := 0; i < 10; i++ {
	// 	num := i + 1
	// 	numbers = append(numbers, num)
	// 	fmt.Println(numbers, cap(numbers))
	// }

	//Flow control
	//Control structures

}
