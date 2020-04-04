package main

import (
        "fmt"
        "math"
        //"interface1/functions"
        //"interface1/others"

)

type shape interface {
  area() float64
  perimeter() float64
}

type circle struct {
  radius float64
}

type rectangle struct {
  side1 float64
  side2 float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2.0
}

func (r rectangle) area() float64 {
	return r.side1 * r.side2
}

func (r rectangle) perimeter() float64 {
	return (r.side1 * 2.0) + (r.side2 * 2.0)
}

func main() {

  var s1 shape = circle{5.0}
  var s2 shape = rectangle{5.0, 10.0}

  fmt.Printf("Shape Type = %T, Shape Value = %v\n", s1, s1)
  fmt.Printf("Area of %T, %f \n", s1, s1.area())
  
  fmt.Printf("Shape Type = %T, Shape Value = %v\n", s2, s2)
  fmt.Printf("Area of %T, %f \n", s2, s2.area())
}
