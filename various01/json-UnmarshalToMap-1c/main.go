package main

//3. Parsing Embedded object in JSON
//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

// Also see
// https://blog.golang.org/json

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Simple Employee JSON  - created by assignment
	empJson := `{
		"id": 101,
		"name": "Some Guy",
		"department": "Cleaning",
		"designation": "Manager",
		"address": {
			"city": "Ranchi",
			"state": "Bihar",
			"country": "India"
		}
	}`

        // a map of "string" key and "interface{}" value
        // NOTE: The input json is stored in this map
        // NOTE: The input json has an embedded json - for which another map is defined after unmarshalling (see below) 
        // interface is used for value so that it can hold any datatype
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	address := result["address"].(map[string]interface{})

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nDesignation :", result["designation"],
		"\nAddress :", address["city"], address["state"], address["country"])
}
