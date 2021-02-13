package main

import (
	"fmt"
)

https://www.google.com/search?q=golang+comparison+operators&oq=golang+compari&aqs=chrome.2.69i57j0i20i263l2j0l7.9913j0j7&sourceid=chrome&ie=UTF-8

func main() {

        /*
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d", &age)
       */
	
	//if boolean_expression {
          //logic or statement here
        // }
          
	/*
         if age < 3 { // 0 to 2
		fmt.Println("infant")

          }else if age >2 && age <13  { //2 to 12
	
		fmt.Println("children")

          }else if age>12 && age <= 19 {

		fmt.Println("teen age")
          }else{
		fmt.Println("adult")
	 }	
        */

        //fixed value
        /*
	switch age {
	case 2:
		fmt.Println("infant")
		fallthrough

	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("children")

	case 13,14,15,16,17,18,19:
		fmt.Println("teen age")

	default:
		fmt.Println("adult")
	}
	*/
        //fmt.Println(age)

        //for loop
         //1,2,3,4,5,6,7,8,9
         // i++ == i=i+1 
         //i=i+1

	/*
	for i:=1; i<=9; i++ {
			
         fmt.Println(i)
        }
        */

	//array string literals
        /*
	students := []string{"Asgor","Mainul","Anonnya"}
        
	for i, std := range students {
  	 fmt.Println(i, std)
	}
	*/
        
        i:=0
	for {
		fmt.Println(i,"hello")
         i++
	}

}
