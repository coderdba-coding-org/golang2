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
    dummyStruct.OthersHelloDummy2()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER\n")

    // dummyStruct := others.DummyStruct{} // already initialized above
    h := handlers.Handler{Others: dummyStruct} //this works for the first field of struct , though does not have the second field
    h  = handlers.Handler{Others: dummyStruct, Handlers: h} // CAUTION - this is convoluted - better not have nesting like this 
                                                            //           with the handler struct itself implementing an interface
                                                            //           Instead, have handler only handle interfaces of one or more others
    fmt.Printf("h = %+v\n", h)
    h.Others.OthersHelloDummy1()  // this line gives - panic: runtime error: invalid memory address or nil pointer dereference
    h.Others.OthersHelloDummy2()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER VIA ANOTHER FUNCTION IN HANDLER\n")
    fmt.Printf("h = %+v\n", h)
    h.HandleOthers()
}
