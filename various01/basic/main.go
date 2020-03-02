package main

import (
        "fmt"
        "basic/pkg/functions"
        "basic/internal/internalfunctions"
)

func main() {  
    functions.Hello()
    internalfunctions.Hello2()
    //fmt.Println("Hello World")

    var x = 111.0
    var ptrX *float64 = &x

    fmt.Printf("The value of x before function call is: %d\n", x)

    functions.TenTimesPtr(ptrX) 
    
    fmt.Printf("The value of x after function call is: %d\n", x)
    
}
