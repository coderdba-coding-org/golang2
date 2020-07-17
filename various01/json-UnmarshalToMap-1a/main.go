package main

//1. Parsing Simple object JSON:
//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

// Also see
// https://blog.golang.org/json

import (
	"encoding/json"
	"fmt"
)

func main() {
	
	// simple json - populate by assignment
	empJson := `{
        "id" : 11,
        "name" : "RamaKrishna",
        "department" : "IT",
        "designation" : "Product Manager"
	}`

        // map of 'string' key and 'interface' value
	// this will allow the 'value' to accept any datatype
	// ALSO - in this case, unmarshal is happening onto a map - not to a struct
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nDesignation :", result["designation"])
}
