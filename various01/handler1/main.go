package main

import (
        "fmt"
        "handler1/handlers"
        "handler1/others"
)

func main() {

    fmt.Printf("\nCALLING FUNCTIONS DIRECTLY\n")
    others.OthersHello()

    fmt.Printf("\nCALLING FUNCTIONS VIA INTERFACE\n")
    dummyStruct := others.DummyStruct{}
    dummyStruct.OthersHelloDummy1()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER\n")
    h := handlers.Handler{}
    fmt.Printf("h = %+v\n", h)
    h.Others.OthersHelloDummy2()
}
