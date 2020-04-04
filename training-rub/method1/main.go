package main

import ("fmt")

type Employee struct{
firstName string
lastName  string
empId int
}

func (e *Employee) GetFirstName() string {
return e.firstName
}

func main() {

// Both with and without & work
//e1 := Employee{"rama", "raghu", 1}
e1 := &Employee{"rama", "raghu", 1}
fmt.Println(e1.GetFirstName())

}
