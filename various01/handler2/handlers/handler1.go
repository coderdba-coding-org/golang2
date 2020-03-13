package handlers

import (
        "fmt"
        "handler2/others"
)

type HandlersInterface interface {
  HandleOthers()
  HandleSelf()
}

type Handler struct {
 Others others.OthersInterface
 Handlers HandlersInterface
 Message string
}

func (i Handler) HandleOthers() {

  fmt.Printf("\nDEBUG: In HandleOthers()\n")
  i.Others.OthersHelloDummy1()
  i.Others.OthersHelloDummy2()
  //i.Others.OthersHello() //this wil fail as this is not there in the interface OthersInterface

}

func (i Handler) HandleSelf() {

  fmt.Printf("\nDEBUG: In HandleSelf()\n")
  fmt.Printf("\nMessage is: %s \n", i.Message)

}
