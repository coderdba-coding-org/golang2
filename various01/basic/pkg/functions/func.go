package functions

import "fmt"

func Hello() {  
    fmt.Println("Hello, really");
}

func TenTimes(a float64) float64 {

return a * 10.0

}

func TenTimesPtr(a *float64) {

*a = *a * 10.0

}
