package main

import (
"io"
"log"
"net/http"
"encoding/json"
)

//json alternative name is optional
type Employee struct{
FirstName string `json:"name"`
LastName  string
EmpID     int
}

func main(){

e1 := Employee {
FirstName: "Rama",
LastName: "Raghu",
EmpID: 1,
}

e2 := Employee {
FirstName: "Krishna",
LastName: "Kanha",
EmpID: 2,
}

employees := []Employee{e1,e2}


// create handlers
helloHandler := func(w http.ResponseWriter, r *http.Request) {
  log.Print("endpoing called")
  io.WriteString(w, "hello world\n")
}

employeesHandler := func(w http.ResponseWriter, r *http.Request) {
  log.Print("employees endpoing called")
  b,_ := json.Marshal(employees) //ignore err - so have "_" instead of err
  io.WriteString(w,string(b))
}

// register the handlers
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/employees", employeesHandler)

// start the server
log.Fatal(http.ListenAndServe(":8888", nil))

}

