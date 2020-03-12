package others

import (
        "fmt"
)

type Others interface {
  OthersHelloDummy1()
  OthersHelloDummy2()
}

type DummyStruct struct {

}

func OthersHello() {
    fmt.Println("Hello, OthersHello");
}

func (d DummyStruct) OthersHelloDummy1() {
    fmt.Println("Hello, OthersHelloDummy 1");
}

func OthersHelloDummy2() {
    fmt.Println("Hello, OthersHelloDummy 2");
}
