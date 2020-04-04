package main

import (
	"fmt"
)

// an empty interface 
// any type will implicitly ending up implementing it because it does not have any function

type EmptyInterface interface{}

// a real interface
type Person interface {
GetFirstName() string
GetLastName() string
GetEmpId() int
}

// a type Employee - and it should also implement all functions of the interface to be an interface
type Employee struct {
firstName string
lastName string
empId int
}

func (e *Employee) GetFirstName() string {
  return e.firstName
}

func (e *Employee) GetLastName() string {
  return e.lastName
}

func (e *Employee) GetEmpId() int {
  return e.empId
}

func main() {

var p Person

// & is required because the functions are for a pointer
// & will not be required if the functions were for an object itself
p = &Employee{"ram", "raghu", 1}
fmt.Println(p)
fmt.Println(p.GetFirstName())
fmt.Println(p.GetLastName())
fmt.Println(p.GetEmpId())

}
