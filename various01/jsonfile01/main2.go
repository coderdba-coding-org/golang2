package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    //"strconv"
)

// Users struct which contains
// an array of users
type NodeList struct {
    Nodes []Node `json:"nodes"`
}

// User struct which contains a name
// a type and a list of social links

type Node struct {
       Id               string `json:"id"`
       Baseline         int    `json:"baseline"`
       TapApp           string `json:"tapApp"`
       MetricSql        string `json:"metricSql"`
       InfluxURL        string `json:"influxURL"`
       DefaultThreshold int    `json:"defaultThreshod"`
       LayerCakeURL     string `json:"layerCakeURL"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("nodeProps.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened nodeProps.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var nodeList NodeList

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &nodeList)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(nodeList.Nodes); i++ {
        fmt.Println("Node Id: " + nodeList.Nodes[i].Id)
    }

}
