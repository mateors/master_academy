package main

import "fmt"
import "reflect"

func main(){

//primitive data types
//rune, byte = alias
//int, float32,string, bool

//Composite Data Types Maps
//array
// bracket = []
// braces/curly braces = {}
// parenthesis = ()

//slice
//more than array

//var students [3]string
//students[0]="Asgor"
//students[1]="Mainul"
//students[2]="Anonnya"
//x := students[0:3]

//x := make([]string, 0)
var fruits []string
fruits = append(fruits, "Apple", "Banana", "Mango")
//fmt.Println(fruits, len(fruits))

//fmt.Printf("%T\n", fruits)
//fmt.Printf("%T\n", students)

b := reflect.TypeOf(fruits).Kind().String()
fmt.Println(b)
 

}