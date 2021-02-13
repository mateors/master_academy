package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello"
	//fmt.Println(str[0:1])
	reverse(str, 0, 4)
	//fmt.Println(Reverse(str))

}
func reverse(str string, start, end int) string {

	var revstr string
	// for i := 0; i < len(str); i++ {
	// 	char := fmt.Sprintf("%c", str[i])
	// 	fmt.Println(i, char)
	// }
	for start < end {
		temp1 := fmt.Sprintf("%c", str[start])
		temp2 := fmt.Sprintf("%c", str[end])
		fmt.Println(start, end, temp1, temp2)
		start++
		end--
	}
	return revstr
}

//ReverseBuilder reverses string using strings.Builder. It's about 3 times faster
//than the one with using a string concatenation
////https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func ReverseBuilder(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

// ReverseByte returns a string with the bytes of s in reverse order.
func ReverseByte(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

//Reverse ..
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

//Reverse2 ..
func Reverse2(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
