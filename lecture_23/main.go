package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// result := Sum(2, 3, 4, 5, 6, 7, 8, 10)
	// fmt.Println(result)

	// strSlc := []interface{}{10, 20, "60"}
	// //result2 := Sum2(strSlc...)
	// result2 := Sum2(10, 20, "60")
	// fmt.Println(result2)

	stooges := [...]string{"Moe", "Larry", "Curly"}
	stooges2 := []string{"Moe", "Larry", "Curly"}
	// fmt.Println(stooges)
	// fmt.Println(stooges2)

	fmt.Printf("%T\n", stooges)
	fmt.Printf("%T\n", stooges2)

	reftype := reflect.TypeOf(stooges).Kind().String()
	reftype2 := reflect.TypeOf(stooges2).Kind().String()
	fmt.Println(reftype)
	fmt.Println(reftype2)

	// for i, slc := range strSlc {
	// 	res := Sum2(slc)
	// 	fmt.Println(i, res)
	// }

}

func Sum2(nums ...interface{}) int64 {

	var res int64
	for _, n := range nums {
		//res += n
		strval := fmt.Sprintf(`%v`, n)
		nint, _ := strconv.ParseInt(strval, 10, 64)
		res += nint

	}
	return res
}

//variadic parameter
//Variadic functions can be called with
//any number of trailing arguments
func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
