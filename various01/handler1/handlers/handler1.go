package handlers

import (
        "fmt"
        "handler1/others"
)

type Handler struct {
 //Message string
 Others others.Others
}

func (i *Handler) HandleOthers() {

  fmt.Printf("\nDEBUG: In HandleOthers()\n")
  //i.Others.OthersHello()

}
