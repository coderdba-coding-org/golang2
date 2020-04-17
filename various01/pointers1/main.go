// training sri

import "fmt"

package main

// does not change the original variable value
func callbyvalue(a, b string) {
a="i am a"
b="i am b"
fmt.Println(a,b)
}

// changes the original variable value
func callbypointer (a,b *string) {

*a = "i am a"
*b = "i am b"

fmt.Println(*a)
fmt.Println(*b)
}


// IF ARRAY/SLICE - changes the original variable value 
// - ALWAYS POINTER COMES IN - NO NEED TO SEND & in the calling function
func callbyvalue2(a []int) {

a[0] = 6

}

func main() {

a:="A"
b:="B"

fmt.Println(a,b)
callbyvalue(a,b)
fmt.Println(a,b)

callbypointer(&a, &b)
fmt.Println(a,b)

// always pointer is sent - no need to send &
array1 := []int{1,2,3}
callbyvalue2(array1)
fmt.Println(array1)

}

