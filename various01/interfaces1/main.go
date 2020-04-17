//training srini
package main

import "fmt"

//func testFunction(a ...interface{}) (interface{}, error) {
//func testFunction(a interface{}) (interface{}, error) {
//func testFunction(a interface{})  {
func testFunction(a ...interface{})  {
  for _, v := range a{
     t := reflect.TypeOf(v)
     fmt.Printf("Type:" #{t}, Kind: #{t.Kind()}, value: #{reflect.ValueOf(v)}\n")
  }
}

func main() {

// THE IDEA HERE IS ANY ITEM IS AN INTERFACE BY DEFAULT
// 'REFLECT' gets the type of the item

testFunction("a", 1) //pass multiple values if ...interface{} is used in the function
fmt.Println()

a:=56
x:=reflect.ValueOf(a).Int()
fmt.Printf("type:#{x} value:#{x}\n)

b:="rama"
y:=reflect.ValueOf(b).String()
fmt.Printf("type:#{y} value:#{y}\n)

}
