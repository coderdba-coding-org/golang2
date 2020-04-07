package main
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	// "strconv"
	// "strings"
)
type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	EmpID     int    `json:"empid"`
}
func main() {
	e1 := Employee{
		FirstName: "John",
		LastName:  "Doe",
		EmpID:     1,
	}
	e2 := Employee{
		FirstName: "bruce",
		LastName:  "wayne",
		EmpID:     2,
	}
	employees := []Employee{e1, e2}
	// 1. create a handler
	employeesHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s : employees endpoint called", r.Method)
		b, _ := json.Marshal(employees)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(b))
	}
	//each employeeHandler
	singleEmployeeHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s : single employee endpoint called", r.Method)
		switch r.Method {
		case http.MethodGet:
			getEmployeeFirstNames(w, r, employees)
		default:
			io.WriteString(w, "Unknown Request")
		}
	}
	// 2. register the handler
        http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/employeefirstnames", singleEmployeeHandler)
	// 3. start the server
	log.Fatal(http.ListenAndServe(":8888", nil))
}
//get employee first names
func getEmployeeFirstNames(w http.ResponseWriter, r *http.Request, employees []Employee) {
  outSliceFirstNames:=[]string{}
  for _, eachEmployee := range employees {
    outSliceFirstNames=append(outSliceFirstNames,eachEmployee.FirstName)
	}
  b, _ := json.Marshal(outSliceFirstNames)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, string(b))
  // io.WriteString(w,outSliceFirstNames)
}
