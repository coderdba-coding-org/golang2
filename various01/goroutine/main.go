package main

import (
        "fmt"
        "time"
)

func func1() {
   for {
        fmt.Println("func1 loop")
        time.Sleep(1 * time.Second)
        //time.Sleep(150 * time.Millisecond)

    }
}

func func2() {
   for {
        fmt.Println("func2 loop")
        time.Sleep(1 * time.Second)
        //time.Sleep(150 * time.Millisecond)
    }
}

func func3() {
   for {
        fmt.Println("func3 loop")
        time.Sleep(1 * time.Second)
        //time.Sleep(150 * time.Millisecond)
    }
}

func main() {
   
   go func1() // 10 min
   //time.Sleep(2 * time.Second)

   go func2() // 30 min (2 min)
   time.Sleep(5 * time.Second)

   // keep something looping in the end - if you want the previous loops to continue
   // otherwise, program puts the previous two in background, REACHES END OF PROGRAM 
   // - and all previous things will automatically be killed of as program has to exit
   //func3() // like a web server

}
