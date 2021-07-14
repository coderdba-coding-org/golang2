// Calculate Reliability number

package influx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

// Node Layer - is a layer of nodes
// Layer in this is a list of strings - each string representing a node
// - it is done so to avoid repeating node details in the layer again
type NodeLayer struct {
	Id    string   `json:"id"`
	Layer []string `json:"layer"`
}

type Entity struct {
	Id      string   `json:"id"`
	Layers  []string `json:"layers"`
	Formula string   `json:"formula"`
}

// InitialDesignEntity - old/initial design - is the main item for which we are calculating reliability
// Layer1 to LayerN are the sets of nodes that make up that layer
type InitialDesignEntity struct {
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

// Calculate reliability
func CalculateReliability(reliabilityFor string) (err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	//entity := Entity{}
	entity := InitialDesignEntity{}

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

	var sloSum float64
	var sloCount float64
	var sloForLayer float64

	// For the layer, get the SLO for each item
	for i := 0; i < len(entity.Layer1); i++ {

		var sloMetValue float64
		var allValue float64
		var slo float64

		fmt.Println("**** Id: ", entity.Layer1[i].Id)

		fmt.Println("**** Running SloMetQuery:")
		resSloMetQuery, err1 := influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].SloMetQuery, entity.Layer1[i].MetricDB)
		if err1 != nil {
			log.Println(err)
			return errors.New("Could not query Slo Met Query ")
		}
		// if no results are returned
		if len(resSloMetQuery[0].Series) == 0 {
			fmt.Printf("WARN: Result of SLO Met Query is null \n")
			sloMetValue = 0
		} else {
			fmt.Println(resSloMetQuery)
			sloMetValue, _ = resSloMetQuery[0].Series[0].Values[0][1].(json.Number).Float64()
			fmt.Printf("sloMetValue from result: %f \n", sloMetValue)
		}

		fmt.Println("**** Running AllQuery:")
		resAllQuery, err2 := influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].AllQuery, entity.Layer1[i].MetricDB)
		if err2 != nil {
			log.Println(err)
			return errors.New("Could not query Slo Met Query ")
		}
		// if no results are returned
		if len(resSloMetQuery[0].Series) == 0 {
			fmt.Printf("WARN: Result of All Query is null \n")
			allValue = 0
		} else {
			fmt.Println(resAllQuery)
			allValue, _ = resAllQuery[0].Series[0].Values[0][1].(json.Number).Float64()
			fmt.Printf("allValue from result: %f \n", allValue)
		}

		if sloMetValue == 0 {
			slo = 0
		} else {
			if allValue == 0 {
				slo = -1
			} else {
				slo = sloMetValue / allValue
			}
		}

		sloSum = sloSum + slo
		sloCount = sloCount + 1

		fmt.Printf("slo calculated for the component: %f \n", slo)
	}

	sloForLayer = sloSum / sloCount
	fmt.Printf("*** slo calculated for the layer: %f \n", sloForLayer)

	return nil
}

// Calculate node's reliability
// Node json format follows Node strut
func CalculateNodeReliability(reliabilityFor string) (sloCalculated float64, err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	entity := Node{}

	// Read the config file for SLO metric definitions
	filePath := "config/node_" + reliabilityFor + ".json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return -200, errors.New("Could not read config file " + filePath)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &entity)

	fmt.Println("Entity is:")
	fmt.Println(entity)

	// PARSE THROUGH THE SQL'S AND GET THE METRIC FROM DB
	fmt.Println("Parsing for entity: " + entity.Id)

	var sloMetValue float64
	var allValue float64
	var slo float64

	fmt.Println("**** Id: ", entity.Id)

	fmt.Println("**** Running SloMetQuery:")
	resSloMetQuery, err1 := influx.QueryInflux(entity.MetricURL, entity.SloMetQuery, entity.MetricDB)
	if err1 != nil {
		log.Println(err)
		return -100, errors.New("Could not query Slo Met Query ")
	}
	// if no results are returned
	if len(resSloMetQuery[0].Series) == 0 {
		fmt.Printf("WARN: Result of SLO Met Query is null \n")
		sloMetValue = 0
	} else {
		fmt.Println(resSloMetQuery)
		sloMetValue, _ = resSloMetQuery[0].Series[0].Values[0][1].(json.Number).Float64()
		fmt.Printf("sloMetValue from result: %f \n", sloMetValue)
	}

	fmt.Println("**** Running AllQuery:")
	resAllQuery, err2 := influx.QueryInflux(entity.MetricURL, entity.AllQuery, entity.MetricDB)
	if err2 != nil {
		log.Println(err)
		return -100, errors.New("Could not query Slo Met Query ")
	}
	// if no results are returned
	if len(resSloMetQuery[0].Series) == 0 {
		fmt.Printf("WARN: Result of All Query is null \n")
		allValue = 0
	} else {
		fmt.Println(resAllQuery)
		allValue, _ = resAllQuery[0].Series[0].Values[0][1].(json.Number).Float64()
		fmt.Printf("allValue from result: %f \n", allValue)
	}

	if sloMetValue == 0 {
		slo = 0
	} else {
		if allValue == 0 {
			slo = -1
		} else {
			slo = sloMetValue / allValue
		}
	}

	fmt.Printf("slo calculated for the component: %f \n", slo)

	return slo, nil
}

// Calculate node-layer's reliability
func CalculateNodeLayerReliability(reliabilityFor string) (sloCalculated float64, err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	entity := NodeLayer{}

	// Read the config file for SLO metric definitions
	filePath := "config/nodelayer_" + reliabilityFor + ".json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return -100, errors.New("Could not read config file " + filePath)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &entity)

	fmt.Println("Entity is:")
	fmt.Println(entity)
	fmt.Println("Parsing for entity: " + entity.Id)

	var sloSum float64
	var sloCount float64
	var sloForLayer float64

	// For the layer, get the node details and calculate SLO for each item
	for i := 0; i < len(entity.Layer); i++ {

		var slo float64

		fmt.Println("**** Layer Node: ", entity.Layer[i])
		slo, err := CalculateNodeReliability(entity.Layer[i])
		if err != nil {
			slo = 0
		}
		sloSum = sloSum + slo
		sloCount = sloCount + 1

		fmt.Printf("slo obtained for the layer-component: %f \n", slo)
	}

	sloForLayer = sloSum / sloCount
	fmt.Printf("*** slo calculated for the layer: %f \n", sloForLayer)

	return sloForLayer, nil
}

// Calculate Entity's reliability
func CalculateEntityReliability(reliabilityFor string) (sloCalculated float64, err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	entity := Entity{}

	// Read the config file for SLO metric definitions
	filePath := "config/entity_" + reliabilityFor + ".json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return -100, errors.New("Could not read config file " + filePath)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &entity)

	fmt.Println("Entity is:")
	fmt.Println(entity)
	fmt.Println("Parsing for entity: " + entity.Id)
	fmt.Println("Formula for entity: " + entity.Formula)

	var sloSum float64
	var sloCount float64
	var sloForEntity float64

	sloSum = 1 // initialize
	formula := entity.Formula

	for i := 0; i < len(entity.Layers); i++ {

		var slo float64

		fmt.Println("**** Layer: ", entity.Layers[i])
		slo, err := CalculateNodeLayerReliability(entity.Layers[i])
		if err != nil {
			slo = 0
		}

		if formula == "product" {
			sloSum = sloSum * slo
		} else {
			sloSum = sloSum + slo
		}

		sloCount = sloCount + 1

		fmt.Printf("slo obtained for the layer: %f \n", slo)
	}

	if formula == "product" {
		sloForEntity = sloSum
	} else {
		sloForEntity = sloSum / sloCount
	}

	fmt.Printf("*** slo calculated for the Entity: %f \n", sloForEntity)

	return sloForEntity, nil
}

// -----------------------------------------------------------------------
// ---- OLDER VERSIONS OF FUNCTIONS ----
// -----------------------------------------------------------------------
// Calculate reliability of a node (old - with Node json was in the format of Entity struct)
func CalculateNodeReliability0(reliabilityFor string) (err error) {

	// READ THE ENTITY FROM ITS CONFIG FILE
	//entity := Entity{}
	entity := InitialDesignEntity{}

	// Read the config file for SLO metric definitions
	filePath := "config/node_" + reliabilityFor + ".json"
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

	var sloSum float64
	var sloCount float64
	var sloForLayer float64

	// For the layer, get the SLO for each item
	for i := 0; i < len(entity.Layer1); i++ {

		var sloMetValue float64
		var allValue float64
		var slo float64

		fmt.Println("**** Id: ", entity.Layer1[i].Id)

		fmt.Println("**** Running SloMetQuery:")
		resSloMetQuery, err1 := influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].SloMetQuery, entity.Layer1[i].MetricDB)
		if err1 != nil {
			log.Println(err)
			return errors.New("Could not query Slo Met Query ")
		}
		// if no results are returned
		if len(resSloMetQuery[0].Series) == 0 {
			fmt.Printf("WARN: Result of SLO Met Query is null \n")
			sloMetValue = 0
		} else {
			fmt.Println(resSloMetQuery)
			sloMetValue, _ = resSloMetQuery[0].Series[0].Values[0][1].(json.Number).Float64()
			fmt.Printf("sloMetValue from result: %f \n", sloMetValue)
		}

		fmt.Println("**** Running AllQuery:")
		resAllQuery, err2 := influx.QueryInflux(entity.Layer1[i].MetricURL, entity.Layer1[i].AllQuery, entity.Layer1[i].MetricDB)
		if err2 != nil {
			log.Println(err)
			return errors.New("Could not query Slo Met Query ")
		}
		// if no results are returned
		if len(resSloMetQuery[0].Series) == 0 {
			fmt.Printf("WARN: Result of All Query is null \n")
			allValue = 0
		} else {
			fmt.Println(resAllQuery)
			allValue, _ = resAllQuery[0].Series[0].Values[0][1].(json.Number).Float64()
			fmt.Printf("allValue from result: %f \n", allValue)
		}

		if sloMetValue == 0 {
			slo = 0
		} else {
			if allValue == 0 {
				slo = -1
			} else {
				slo = sloMetValue / allValue
			}
		}

		sloSum = sloSum + slo
		sloCount = sloCount + 1

		fmt.Printf("slo calculated for the component: %f \n", slo)
	}

	sloForLayer = sloSum / sloCount
	fmt.Printf("*** slo calculated for the layer: %f \n", sloForLayer)

	return nil
}
