package main

import (
        "fmt"
        "time"
        "infiniteloop/pkg/functions"
        "infiniteloop/internal/internalfunctions"
)

func main() {  
    functions.Hello()
    internalfunctions.Hello2()
    //fmt.Println("Hello World")

    // Infinite loop
    for {

      var x = 111.0
      var ptrX *float64 = &x

      fmt.Printf("The value of x before function call is: %d\n", x)
      functions.TenTimesPtr(ptrX) 
      fmt.Printf("The value of x after function call is: %d\n", x)

      // both these time functions work
      //time.Sleep(1000 * time.Millisecond)
      time.Sleep(time.Duration(1) * time.Second)
    
    }
}
