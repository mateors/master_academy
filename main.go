package main

import "fmt"

func main(){

//primitive data types
//rune, byte = alias
//int, float32,string, bool

//Composite Data Types Maps
//Maps
//Key = value

//null
//nil

//var x map[string]string
x := make(map[string]string)

x["name"] = "Mostain"
x["height"] = "5.7"
x["address"] = "Dhaka"

//delete(x, "height")

fmt.Println(x)

}