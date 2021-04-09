package main

import (
	"fmt"

	"github.com/Mahmud139/vatcalculator"
)

func main() {

	vatAmount := vatcalculator.InclusiveTax(100, 15)
	fmt.Println(vatAmount)
}
