package main

import "fmt"

func update(a *int){
 fmt.Println(a)
 *a=*a+10
}

func main(){
 
 var x int
 var y *int

 fmt.Println("x value is", x)
 fmt.Println("x memory address is", &x)

 fmt.Println("y value is", y)
 fmt.Println("y memory address is", &y)

 x=10 //assignment
 y=&x //referencing
 fmt.Println("x is", x) //accessing
 fmt.Println("y is", y) 

 fmt.Println("y dererencing value is", *y)

 update(&x)

 fmt.Println(x)

 

}

//https://tour.golang.org/moretypes/1