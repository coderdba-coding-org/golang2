package main

import (
        "fmt"
        "handler2/others"
        "handler2/handlers"

)

//type Handler struct {
 //Message string
 //Others others.Others
//}

func main() {

    fmt.Printf("\nCALLING FUNCTIONS DIRECTLY\n")
    others.OthersHello()

    fmt.Printf("\nCALLING FUNCTIONS VIA INTERFACE\n")
    dummyStruct := others.DummyStruct{"dummyvalue"}  // initialize structure that implements the interface in others package
    dummyStruct.OthersHelloDummy1()  // call the implemented function
    dummyStruct.OthersHelloDummy2()  // call the implemented function

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER\n")

    // dummyStruct := others.DummyStruct{} // already initialized above
    //h := handlers.Handler{Others: dummyStruct, Message: "Initial"} //this works for the first field of struct , though does not have the second field
    h := handlers.Handler{}
    h  = handlers.Handler{Others: dummyStruct, Handlers: h, Message: "Initial-Message"} // CAUTION - this is convoluted - better not have nesting like this
                                                            //           with the handler struct itself implementing an interface
                                                            //           Instead, have handler only handle interfaces of one or more others
    //h := handlers.Handler{Others: {}, Handlers: {}, Message: "InitialMessage"}
    fmt.Printf("h = %+v\n", h)
    h.Others.OthersHelloDummy1()
    h.Others.OthersHelloDummy2()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER VIA ANOTHER FUNCTION IN HANDLER\n")
    fmt.Printf("h = %+v\n", h)
    h.HandleOthers()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER VIA A CALLER CALLING ANOTHER FUNCTION IN HANDLER\n")
    fmt.Printf("h = %+v\n", h)

    fmt.Printf("\nThis function has input as handler\n")
    h.CallerHandleOthers1()

    fmt.Printf("\nThis function has input as pointer to handler\n")
    h.CallerHandleOthers2()

    fmt.Printf("\nCALLING FUNCTIONS VIA HANDLER OF A FUNCTION IN HANDLER OF HANDLER\n")
    fmt.Printf("h = %+v\n", h)
    h.HandleSelf()
}
