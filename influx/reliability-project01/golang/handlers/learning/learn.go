// Calculate Reliability number

package influx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	client "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"

	// local packages
	influx "dbtrends/handlers/influx"
)

var rel1_config string = "../../config/reliability_sql_db.json"

// Node is each influencer in a "layer"
type Node struct {
	Id           string  `json:"id"`
	SloMetQuery  string  `json:"sloMetQuery"`
	AllQuery     string  `json:"allQuery"`
	MetricURL    string  `json:"metricURL"`
	MetricDB     string  `json:"metricDB"`
	MetricDBType string  `json:"metricDBType"`
	Baseline     float64 `json:"baseline"`
}

// Entity is the main item for which we are calculating reliability
// Layer1 to LayerN are the sets of nodes that make up that layer
type Entity struct {
	Entity string `json:"entity"`
	Layer1 []Node `json:"layer1"`
}

// Reliability of each node in a "layer"
type NodeReliability struct {
	Id             string  `json:"id"`
	SloMetValue    float64 `json:"sloMetValue"`
	AllValue       float64 `json:"allValue"`
	ReliabilityPct float64 `json:"reliabilityPct"`
}

// Reliability of an entity
type EntityReliability struct {
	Entity            string            `json:"entity"`
	Layer1Reliability []NodeReliability `json:"layer1"`
}

// Read the master-list of nodes into memory
func SampleReliability(reliabilityFor string) (res []client.Result, err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	//var entity Entity
	entity := Entity{}

	// Read the config file for SLO metric definitions
	filePath := "config/reliability_" + reliabilityFor + ".json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return errors.New("Could not read config file " + filePath)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &entity)

	fmt.Println("Entity is:")
	fmt.Println(entity)

	// PARSE THROUGH THE SQL'S AND GET THE METRIC FROM DB
	fmt.Println("Parsing for entity: " + entity.Entity)

	var res1 []client.Result
	var res2 []client.Result

	for i := 0; i < len(entity.Layer1); i++ {
		fmt.Println("**** Id: ", entity.Layer1[i].Id)
		fmt.Println("**** Running SloMetQuery:")
		res1, err := influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].SloMetQuery, entity.Layer1[i].MetricDB)
		//fmt.Println("**** Running AllQuery:")
		//res2, err = influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].AllQuery, entity.Layer1[i].MetricDB)
	}
	return res1, nil
}
