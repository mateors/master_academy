package main

import "fmt"

type Book struct {
	Name  string
	Price int
}

type Human interface {
	Speak()
}

type Doctor struct {
	Name   string
	Degree string
	Age    int
}

type Engineer struct {
	Name   string
	Degree string
	Age    int
}

func (d Doctor) Speak() {
	fmt.Println(d.Name, "can speak")
}

func (d Doctor) Surgery() {
	fmt.Println(d.Name, "can surgery")
}

func (e Engineer) Speak() {
	fmt.Println(e.Name, "can speak")
}

func (e Engineer) Programming() {
	fmt.Println(e.Name, "can code")
}

func main() {

	d := Doctor{"Sanzida", "MBBS", 28}
	//d.Speak()

	e := Engineer{Name: "Hasan", Degree: "BSC", Age: 30}
	//e.Speak()

	humans := [2]Human{d, e}
	//fmt.Println(humans)
	for _, v := range humans {

		//fmt.Println(i,v)
		switch v.(type) {
		case Doctor:
			dd := v.(Doctor)
			dd.Surgery()
		case Engineer:
			v.(Engineer).Programming()

		}
	}

}
