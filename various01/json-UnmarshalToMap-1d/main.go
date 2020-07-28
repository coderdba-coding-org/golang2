package main

//4. Parsing Embedded object in Array of JSON
//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

// Also see
// https://blog.golang.org/json

import (
	"encoding/json"
	"fmt"
)

func main() {
	//Simple Employee JSON object which we will parse
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 101,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 102,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empArray), &results)

	for key, result := range results {
		address := result["address"].(map[string]interface{})
		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])
		fmt.Println("Address :", address["city"], address["state"], address["country"])
	}
}
