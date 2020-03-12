package main

import (
        "fmt"
        "funcstruct/pkg/functions"
        "funcstruct/internal/internalfunctions"
)

//type Handlers struct {
//	Func         functions.Func
//	FuncInternal internalfunctions.FuncInternal
//}

func main() {

    fmt.Printf("\nCALLING FUNCTIONS DIRECTLY\n")
    functions.Hello()
    internalfunctions.Hello2()
    //fmt.Println("Hello World")

    var x = 111.0
    var ptrX *float64 = &x

    fmt.Printf("The value of x before function call is: %d\n", x)

    functions.TenTimesPtr(ptrX)
    fmt.Printf("The value of x after function call is: %d\n", x)

    //
    //

    fmt.Printf("\n\nCALLING FUNCTIONS VIA HANDLER INTERFACE\n")

    h := functions.FuncHandler{}
    h.Init()

    //f := functions.Func{}
    //fi := internalfunctions.FuncInternal{}

    var y = 222.0
    var ptrY *float64 = &y

    fmt.Printf("The value of y before function call is: %d\n", y)

    functions.TenTimesPtr(ptrY)
    fmt.Printf("The value of y after function call is: %d\n", y)

    h.HandlerTenTimesPtr(ptrY)
    fmt.Printf("The value of y after function call - with handler - is: %s %d\n", h.Message, y)

    h.ModMessage()
    h.HandlerTenTimesPtr(ptrY)
    fmt.Printf("The value of y after function call - with handler - is: %s %d\n", h.Message, y)

    h.HandlerTenTimesPtr(ptrY)
    fmt.Printf("The value of y after function call - with handler - is: %s %d\n", h.Message, y)

    // this is not working if hello2 is uncommented inside func.go
    h.HandlerHello()

    // this is not working
    //h.FunctionsInternal.Hello2()

}
