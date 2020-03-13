package others

import (
        "fmt"
)

type OthersInterface interface {
  OthersHelloDummy1()
  OthersHelloDummy2()
}

type DummyStruct struct {
  Dummy string
}

func OthersHello() {
    fmt.Println("Hello, OthersHello");
}

func (d DummyStruct) OthersHelloDummy1() {
    fmt.Println("Hello, OthersHelloDummy 1");
}

func (d DummyStruct) OthersHelloDummy2() {
    fmt.Println("Hello, OthersHelloDummy 2");
}
