package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// var n = 101

	// if n%10 == 0 {
	// 	fmt.Println("works")
	// } else {
	// 	fmt.Println(n % 5)
	// }

	// isPD := isPalindrome(101)
	// fmt.Println(isPD)

	// 	Example test:   (['sander', 'amy', 'ann', 'michael'], ['123456789', '234567890', '789123456', '123123123'], '1')
	// WRONG ANSWER (got sander expected ann)

	// Example test:   (['adam', 'eva', 'leo'], ['121212121', '111111111', '444555666'], '112')
	// OK

	// var A = []string{"pim", "pom"}
	// var B = []string{"999999999", "777888999"}

	// var A = []string{"sander", "amy", "ann", "michael"}
	// var B = []string{"123456789", "234567890", "789123456", "123123123"}
	// var P = "1"
	// r := Solution(A, B, P)
	// fmt.Println(r)

	//animals := []string{"snail", "dog", "cow", "elephant", "chicken", "mouse"}
	animals := []string{"sander", "amy", "ann", "michael"}
	fmt.Println(animals)

	sort.Strings(animals)
	fmt.Println(animals)

	sort.Slice(animals, func(i, j int) bool {
		return len(animals[i]) < len(animals[j])
	})
	fmt.Println(animals)

}

func Solution(A []string, B []string, P string) string {

	sort.Strings(A)
	//sort.Strings(B)
	//sort.Strings()

	for indx, phone := range B {
		if strings.Contains(phone, P) {

			sort.Slice(A, func(i, j int) bool {
				return len(A[i]) < len(A[j])
			})

			return A[indx]
		}
	}
	return "NO CONTACT"
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	t := x
	var r int
	for x > 0 {
		r = r*10 + x%10
		x /= 10
		fmt.Println(r, x)
	}
	return t == r
}
