package handlers

import (
        "fmt"
)

func (j Handler) CallerHandleOthers2() {
  fmt.Printf("\nDEBUG: In CallerHandleOthers()\n")
  //i.Others.OthersHello() //this wil fail as this is not there in the interface OthersInterface
  j.Others.OthersHelloDummy1()
  j.Others.OthersHelloDummy2()
}

func (j *Handler) CallerHandleOthers1() {
  fmt.Printf("\nDEBUG: In CallerHandleOthers()\n")
  //i.Others.OthersHello() //this wil fail as this is not there in the interface OthersInterface
  j.Others.OthersHelloDummy1()
  j.Others.OthersHelloDummy2()
}
