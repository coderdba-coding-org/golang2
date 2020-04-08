package main

//https://motyar.github.io/golang-pretty-print-struct/

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {

	var p Person
	p.name = "Rama"
	p.age = 20

        fmt.Printf("%s\n", "Printing the whole structure - without type")
	fmt.Printf("%+v\n\n", p) //With name and value

        fmt.Printf("%s\n", "Printing the whole structure - with type")
	fmt.Printf("%#v\n\n", p) //with name, value and type

        fmt.Printf("%s\n", "Printing only Name")
	fmt.Printf("%v\n\n", p.name)

        fmt.Printf("%s\n", "Printing only age")
	fmt.Printf("%v\n\n", p.age)
}

