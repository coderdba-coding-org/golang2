package internalfunctions

import "fmt"

type FuncInternal interface {
  Hello2()
}

func Hello2() {
    fmt.Println("Hello2, really");
}
