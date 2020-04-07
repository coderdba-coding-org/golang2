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

type EmployeeFirstName struct{
   FirstName string `json:"name"`
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

// without pointer
employeesFirstNames := []EmployeeFirstName{}

// with pointer
//employeesFirstNames := &[]EmployeeFirstName{}

// create handlers
helloHandler := func(w http.ResponseWriter, r *http.Request) {
  log.Print("endpoing called")
  io.WriteString(w, "hello world\n")
}

employeesHandler := func(w http.ResponseWriter, r *http.Request) {
  log.Print("employees endpoing called")
  b,_ := json.Marshal(employees) //ignore err - so have "_" instead of err
  w.Header().Set("Content-Type", "application/json")
  io.WriteString(w,string(b))
}

employeesFirstNameHandler := func(w http.ResponseWriter, r *http.Request) {

  log.Print("employees firstname endpoing called")

  // without pointer
  employeesFirstNames = nil 

  // with pointer
  //*employeesFirstNames = nil 

  efn := EmployeeFirstName{}

  //i := 0

  for _, employee := range employees {
      
      efn = EmployeeFirstName{FirstName: employee.FirstName}
      // This is working - without pointer reference
      employeesFirstNames = append(employeesFirstNames, efn)
    
      // This is working - with pointer reference
      //*employeesFirstNames=append((*employeesFirstNames),efn)
      //i++
  }

  b,_ := json.Marshal(employeesFirstNames) //ignore err - so have "_" instead of err
  w.Header().Set("Content-Type", "application/json")
  io.WriteString(w,string(b))
}

// register the handlers
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/employees", employeesHandler)
http.HandleFunc("/employeefirstnames", employeesFirstNameHandler)

// start the server
log.Fatal(http.ListenAndServe(":8888", nil))

}
