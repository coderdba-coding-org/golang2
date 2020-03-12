package functions

import (
        "fmt"
        //"funcstruct/internal/internalfunctions"
)

type Func interface {
  Hello()
  TenTimes(a float64)
  TenTimesPtr(a *float64)
  New()
}

type FuncHandler struct {
	Message string
  Functions Func  // reference the interface here
  //FunctionsInternal internalfunctions.FuncInternal  // reference the interface here
}

func Hello() {
    fmt.Println("Hello, really");

}

func (i *FuncHandler) HandlerHello() {
    fmt.Println("Hello in HandlerHello, really");
    fmt.Printf("DEBUG: %s - is the message in the handler\n", i.Message)
    i.Functions.Hello()
    //i.FunctionsInternal.Hello2()
}

func TenTimes(a float64) float64 {

return a * 10.0

}

func TenTimesPtr(a *float64) {

*a = *a * 10.0

}

func (i *FuncHandler) Init() {

i.Message = "Func Handler says: "

}

func (i *FuncHandler) HandlerTenTimesPtr(a *float64) {

*a = *a * 10.0

}

func (i *FuncHandler) ModMessage() {

i.Message = "Func Handler now says:"

}
