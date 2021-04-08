package main
//https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

import (
        "encoding/json"
        "fmt"
)

func main() {

/* Structure of an Influx result single line */
resultJson := `[
  {
    "statement_id": 0,
    "Series": [
      {
        "name": "gom-api-response-time",
        "columns": [
          "time",
          "value"
        ],
        "values": [
          [
            "2021-04-08T06:03:12.190018436Z",
            113282624.05424652
          ]
        ]
      }
    ],
    "Messages": "Hi There"
  },
  {
    "statement_id": 0,
    "Series": [
      {
        "name": "gom-api-response-time",
        "columns": [
          "time",
          "value"
        ],
        "values": [
          [
            "2021-04-08T06:03:12.190018436Z",
            113282624.05424652
          ],
          [
            "4022-04-08T06:03:12.190018436Z",
            413282624.05424652
          ]
        ]
      }
    ],
    "Messages": "Hi There"
  }
]`

// Variable for Unmarshalled resultJson
//  [] for array, of map of [string] key, and interface{} value
var results []map[string]interface{}


fmt.Println("Printing result variable (before marshaling):")
fmt.Println(resultJson)

json.Unmarshal([]byte(resultJson), &results)

for key, result := range results {
        fmt.Println("Reading Value for Key :", key)

        id := result["statement_id"]
        fmt.Println("statement_id is :", id)

        //series := result["Series"] //works
        series := result["Series"].([]interface{}) //works
        //series := result["Series"].([]map[string]interface{}) //does not work
        fmt.Println("Series is :", series)

        messages := result["Messages"]
        fmt.Println("Messages is :", messages)

        //values := series["values"]
        //values := series[0]["values"]
        //values := series[0] // this much works
        //values := series[0].(map[string]interface{}) // this much works
        //values := series[0].(map[string]interface{})["values"]  // this much works
        //values := series[0].(map[string]interface{})["values"].([]interface{}) //this much works
        values := series[0].(map[string]interface{})["values"].([]interface{})
        fmt.Println("values is :", values)
        fmt.Println("values[0] is :", values[0])
        fmt.Println("values[0][1] is :", values[0].([]interface{})[1])

        //fmt.Println("metric date is :", values[0])
        //fmt.Println("metric value is :", values[1])

        //Reading each value by its key
        //fmt.Println("Id :", result["id"],
                //"- Name :", result["name"],
                //"- Department :", result["department"],
                //"- Designation :", result["designation"])
        //fmt.Println("Address :", address["city"], address["state"], address["country"])
}

} //end of main
