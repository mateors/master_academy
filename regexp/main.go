package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	// regxp, err := regexp.Compile(`a`)
	// if err != nil {

	// 	fmt.Println(err.Error())
	// 	return
	// }
	//resa := regxp.FindAllStringSubmatch(txtdata, -1)

	res := isMatch(`abbdbb`, "ab*d")
	fmt.Println(res)

}

func isMatch(text string, pattern string) bool {

	if strings.ContainsAny(pattern, "*.") == true {

		rp, err := regexp.Compile("ab*d")
		if err != nil {

			return false
		}
		return rp.MatchString(text)

		// isMatch, err := regexp.MatchString(pattern, text)
		// if err != nil {
		// 	return false
		// }
		// return isMatch
	}
	if text == pattern {
		return true
	}
	return false
}
