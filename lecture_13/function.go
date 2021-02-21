package main

import "fmt"

//example-1
/*
func add(x int, y int)int{
//body
r := x+y
return r
}
*/

/*
//example-2
func add(x, y int)int{
//body
r := x+y
return r
}
*/

//example-3, Named return values

/*
func add(x, y int) (r int){
//body
 r= x+y
 return r
}
*/


//example-4
func add(x, y int) (r int){
//body
 r= x+y
 return
}

func Add(x, y int) (r int){
//body
 r= x*y
 return
}

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
	 // Return statement without specify variable name
}

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

func main(){

 //we arite code here
 //x:=Add(10, 30)
 //fmt.Println(x)
 //a,p := rectangle(10, 10)
 //fmt.Println(a, p)
	
  /*
  number:=10
  name:="Mostain"
  update(&number, &name)
  fmt.Println(number, name)
*/

 a := func(x, y int)(r int){
 	r=x*y
 	return
    }(10,10)

 //fmt.Println(a(10,10))
fmt.Println(a)

}

