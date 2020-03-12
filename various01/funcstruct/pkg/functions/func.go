package functions

import "fmt"

type Func interface {
  Hello()
  TenTimes(a float64)
  TenTimesPtr(a *float64)
}

type FuncHandler struct {
	Message string
}

func Hello() {
    fmt.Println("Hello, really");
}

func TenTimes(a float64) float64 {

return a * 10.0

}

func TenTimesPtr(a *float64) {

*a = *a * 10.0

}

func (h *FuncHandler) Init() {

h.Message = "Func Handler says: "

}

func (h *FuncHandler) TenTimesPtr(a *float64) {

*a = *a * 10.0

}
