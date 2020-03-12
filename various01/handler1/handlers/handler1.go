package handlers

import (
        "fmt"
        "handler1/others"
)

type HandlersInterface interface {
  HandleOthers()
}

type Handler struct {
 //Message string
 Others others.OthersInterface
 Handlers HandlersInterface
}

func (i *Handler  ) HandleOthers() {

  fmt.Printf("\nDEBUG: In HandleOthers()\n")
  //i.Others.OthersHello()

}
