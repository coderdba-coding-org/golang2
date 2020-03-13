package handlers

import (
        "fmt"
        "handler1/others"
)

type HandlersInterface interface {
  HandleOthers()
}

type Handler struct {
 Others others.OthersInterface
 Handlers HandlersInterface
 //Message string
}

func (i Handler) HandleOthers() {

  fmt.Printf("\nDEBUG: In HandleOthers()\n")
  //i.Others.OthersHello() //this wil fail as this is not there in the interface OthersInterface
  i.Others.OthersHelloDummy1()
  i.Others.OthersHelloDummy2()
}
