package main

import (
	"fmt"
	"reflect"
)

type human interface {
	speak()
}

//type Degree = string //alis of string
type Degree string //custom type

type doctor struct {
	name      string
	age       int
	education Degree
}

type engineer struct {
	name      string
	age       int
	education Degree
}

func (d doctor) surgery() {

	fmt.Println(d.name + " can surgery")

}

func (d doctor) speak() {

	fmt.Println(d.name + " can speak")

}

func (e engineer) programming() {

	fmt.Println(e.name + " can make software")

}
func (e engineer) speak() {

	fmt.Println(e.name + " can speak")

}

//you can ...
func greet(h human) {

	//polymorphism == বহুরূপতা
	//Poly = more
	//Morph = change
	//many different types
	//d := h.(doctor)
	//e := h.(engineer)
	//fmt.Println(d.name, d.education)
	//fmt.Println(e.name, e.education)
	//assertion = জোর / দাবি করা / kind of certify / attested / সত্যায়িত
	//convert = রূপান্তর =  pertrol to cng
	//Assertions are statements used to test assumptions made by programmer
	switch h.(type) {

	case doctor:
		fmt.Println(h.(doctor).education)

	case engineer:
		fmt.Println(h.(engineer).education)
	}

	fmt.Println("we can greet to ", h)
}

//you can offer food

func main() {

	d1 := doctor{name: "Sanzida", age: 28, education: "MBBS"}
	e1 := engineer{name: "Mostain", age: 37, education: "CSE"}

	//fmt.Printf("d1 is %T\n", d1)
	//fmt.Printf("e1 is %T\n", e1)
	humans := []human{d1, e1}
	d1UnderLyingType := reflect.TypeOf(d1.education).Kind().String()
	fmt.Printf("d1 is %T\n", d1.education)
	fmt.Println(d1UnderLyingType)
	fmt.Println(d1, e1, humans)

	greet(humans[0])

}
