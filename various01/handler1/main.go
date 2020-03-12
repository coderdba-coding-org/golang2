package main

import (
        "fmt"
        //"handler1/handlers"
        "handler1/others"
)

//type Handlers struct {
//	Func         functions.Func
//	FuncInternal internalfunctions.FuncInternal
//}

func main() {

    fmt.Printf("\nCALLING FUNCTIONS DIRECTLY\n")
    others.Others1Hello()

}
