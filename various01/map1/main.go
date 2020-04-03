package main

import (
	"fmt"
)

func main() {

//m := make(map[int]string)

m := map[int]string{
0: "hello",
1: "ola",
}

fmt.Println(m)

for k,v := range m {
  fmt.Printf("key = %d, val = %s \n", k, v)
}

}
