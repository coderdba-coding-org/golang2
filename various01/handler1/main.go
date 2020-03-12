package main

import (
        "fmt"
        "handler1/others"
        "handler1/handlers"

)

//type Handler struct {
 //Message string
 //Others others.Others
//}

func main() {

    fmt.Printf("\nCALLING FUNCTIONS DIRECTLY\n")
    others.OthersHello()

    fmt.Printf("\nCALLING FUNCTIONS VIA INTERFACE\n")
    dummyStruct := others.DummyStruct{}
    dummyStruct.OthersHelloDummy1()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER\n")
    h := handlers.Handler{}
    //h := Handler{}
    fmt.Printf("h = %+v\n", h)
    h.Others.OthersHelloDummy2()  // this line gives - panic: runtime error: invalid memory address or nil pointer dereference
}
