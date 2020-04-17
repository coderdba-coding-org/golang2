//training sri

package main

import ("fmt"
"os"
)

func divide(a, b *float64) (float64m, error) {
if b == 0 {
// use one of the below
// however, 0 is not the error code - it is just a dummy resut returned
return 0, errors.New('cant divide by zero')
return 0, fmt.Error('cant divide by zero')
} 

// if b is not 0, then return the value of a divided by b and nil error
return a/b, nil
 
}

func main() {

// check which syntax is correct
_,err := os.Open("temp.txt")
_,err := os.Open(name: "temp.txt")
if err != nil {
panic(err)
}

result, err := divide(1.31,0)
if err != nil{
fmt.Println ("Error: ", err.Error()
} else {
fmt.Println ("Result: ", result)
}

}

