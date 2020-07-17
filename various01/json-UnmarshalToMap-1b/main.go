package main

//2. Parsing Simple Array JSON:k
//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

// Also see
// https://blog.golang.org/json

import (
	"encoding/json"
	"fmt"
)

func main() {

	// slice/array of jsons - created by assignment
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director"
		},
		{
			"id": 11,
			"name": "RamaKrishna",
			"department": "IT",
			"designation": "Product Manager"
		},
		{
			"id": 12,
			"name": "Seetha",
			"department": "IT",
			"designation": "Team Lead"
		}
	]`

	// a slice "[]" of maps - with map of "string" key and "interface{}" value
        // NOTE: Individual jsons in the input slice/array is stored in multiple maps of key-value pairs
        //       - each map represents one json 
	// interface is used for value so that it can hold any datatype
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empArray), &results)

	for key, result := range results {

		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])
	}
}
