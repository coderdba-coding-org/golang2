package main

import (
        "github.com/gin-gonic/gin"
	//"errors"
	"log"
        //"encoding/json"
	"time"
        "net/http"
	"encoding/json"
	"fmt"
        "os"
	"io/ioutil"
        //"strconv"

        //client "github.com/influxdata/influxdb/client/v2"
        client "github.com/influxdata/influxdb1-client/v2"
)

type OracleMetric0 struct {
	time        time.Time 
	DbName      int 
	MetricValue int 
}
type OracleMetric struct {
	Time        time.Time 
	DbName      int 
	MetricValue int 
}

// List of nodes
type NodeList struct {
       Nodes []Node `json:"nodes"`
}

// Struct for node properties
type Node struct {
       Id               string `json:"id"`
       Baseline         int    `json:"baseline"`
       TapApp           string `json:"tapApp"`
       MetricSql        string `json:"metricSql"`
       InfluxURL        string `json:"influxURL"`
       DefaultThreshold int    `json:"defaultThreshod"`
       LayerCakeURL     string `json:"layerCakeURL"`
}

// List of input nodes
type NodeInputList struct {
     Nodes []NodeInput `json:"nodes"`
}

// Struct for input node properties
type NodeInput struct{
      Id	string `json:"id"`
      Group	int `json:"group"`
      Baseline	int `json:"baseline"`
      Value	int `json:"value"`
}

var nodeList NodeList

//---------------//

func main() {

	log.Println("entering main()")
        
        /*
        // Read Node Properties
        file, _ := ioutil.ReadFile("./nodeProps.json")
	log.Println("read the file in main()")
	nodeList := NodeList{}
	_ = json.Unmarshal([]byte(file), &nodeList)
	log.Println("nodelist length: ", len(nodeList.nodes))
        */


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
    
        // we initialize our nodeList array
        //var nodeList NodeList
    
        // we unmarshal our byteArray which contains our
        // jsonFile's content into 'NodesList' which we defined above
        json.Unmarshal(byteValue, &nodeList)
    
        // we iterate through every Node within our NodeList array and
        // print out the user Type, their name, and their facebook url
        // as just an example
        fmt.Println(nodeList)
        for i := 0; i < len(nodeList.Nodes); i++ {
            fmt.Println("Node Id: " + nodeList.Nodes[i].Id)
        }

        // Web server
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
                     "message": "Welcome!",
		})
	})

        // health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
		})
	})


	r.GET("/nodes", func(c *gin.Context) {

                // works
                c.ShouldBind(&nodeList) //works
		c.IndentedJSON(http.StatusOK, nodeList) //works
		//c.JSON(http.StatusOK, nodeList) //works

                /* does not work
                if err := c.BindJSON(&nodeList); err != nil {
                        log.Println("JSON Binding error")
                        log.Println(err)
                        log.Println(nodeList)
			return
		}
                */

                // does not work
                //b, _ := json.Marshal(nodeList);
                //c.JSON(200, b)

                // does not work
                //b, err := json.Marshal(nodeList);
		//c.JSON(200, gin.H{
                     //b,
		//})
	})


        // db query endpoint - runs a pre-set query
	//r.GET("/query", QueryDbLocalMac)
	r.GET("/query", QueryDbTgt)

        // get the metric value for a node
	r.GET("/getmetric/:nodeName", GetNodeMetric)

        // get the metric value for input node list
	r.GET("/getnodelistmetric/", GetNodeInputMetric)

        //--------------------------
        /* does not work
	//r.GET("/query", gin.H(queryDb))
	//r.GET("/query", func(c *gin.Context) { queryDb })
        */
        //--------------------------



        // original port 8080
	//err := r.Run("0.0.0.0:8080")
	err = r.Run("0.0.0.0:8080")
	//err := r.Run("0.0.0.0:5050")

	if err != nil {
	}

}

//func GetValueFromResult(resultJson []client.Result) {
func GetValueFromResult(resultJson []byte) {

   var results []map[string]interface{}

   fmt.Println("Printing result variable (before marshaling):")
   fmt.Println(resultJson)

   json.Unmarshal([]byte(resultJson), &results)
   //json.Unmarshal(resultJson, &results)

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

   }
}

func QueryDbTgt(c *gin.Context) {
        //cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"
        cmd := "SELECT count(metric_value) FROM oracle_dbstatus WHERE _blossom_id = 'CI02989373' AND application = 'RDBMS' AND oracledb_service = 'ORDPROD' AND metric_name = 'instance_status' AND time > now() - 30m"
        log.Println("Query Is", cmd)

	q := client.Query{
		Command:  cmd,
		Database: "metrics",
	}

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{
		//Addr: "http://localhost:8086",
		//Addr: "https://metrics-shared.prod.company.com/",
		Addr: "https://metricsengine-shared.prod.company.com/",
              
	})
        if err != nil {
           log.Println("Influx DB Client error")
           c.AbortWithStatus(http.StatusNotFound)
        }

	if response, err := dbclient.Query(q); err == nil {
		if response.Error() != nil {
                    log.Println("Error in query results")
                    log.Println(response)
                    c.AbortWithStatus(http.StatusNotFound)
		}
		res := response.Results
                //log.Println("got query results")
                //log.Println(res)
		c.JSON(http.StatusOK, res)
	} else {
                log.Println("failed to get query results")
                log.Println(response)
                c.AbortWithStatus(http.StatusNotFound)
	}
}

func GetNodeMetric(c *gin.Context) {

    nodeName := c.Params.ByName("nodeName")
    var metricSQL string
    var influxURL string
    var baseLine  int
    var defaultThreshold int

    for i := range nodeList.Nodes {
        if nodeList.Nodes[i].Id == nodeName {
            metricSQL = nodeList.Nodes[i].MetricSql
            influxURL = nodeList.Nodes[i].InfluxURL
            baseLine = nodeList.Nodes[i].Baseline
            defaultThreshold = nodeList.Nodes[i].DefaultThreshold
          
            log.Printf(nodeName + " found" )
            log.Printf(metricSQL)
            log.Printf(influxURL)
            log.Printf("%d", baseLine)
            log.Printf("%d", defaultThreshold)
 
            break
        }
    }


	q := client.Query{
		Command:  metricSQL,
		Database: "metrics",
	}

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{
		//Addr: "http://localhost:8086",
		//Addr: "https://metrics-shared.prod.company.com/",
		//Addr: "https://metricsengine-shared.prod.company.com/",
                Addr: influxURL,
	})
        if err != nil {
           log.Println("Influx DB Client error")
           c.AbortWithStatus(http.StatusNotFound)
        }

	if response, err := dbclient.Query(q); err == nil {
		if response.Error() != nil {
                    log.Println("Error in query results")
                    log.Println(response)
                    c.AbortWithStatus(http.StatusNotFound)
		}
		res := response.Results
                log.Println("got query results")
                log.Println(res)
		c.JSON(http.StatusOK, res)
	} else {
                log.Println("failed to get query results")
                log.Println(response)
                c.AbortWithStatus(http.StatusNotFound)
	}
     
    //c.JSON(http.StatusOK,gin.H{ "message": "in GetNodeMetric!", } )

}

func GetNodeInputMetric(c *gin.Context) {

    //variables for node master properties
    var metricSQL string
    var influxURL string
    var baseLine  int
    //var defaultThreshold int

    // Read the input node list
    jsonFile, err := os.Open("nodeInput.json")
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println("Successfully Opened nodeProps.json")

    defer jsonFile.Close()
    
    // read our opened file as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)
 
    // variables for input node list, and a node from it
    var nodeInputList NodeInputList
    var nodeName      string
 
    // jsonFile's content into 'NodesInputList' which we defined above
    json.Unmarshal(byteValue, &nodeInputList)

    log.Println("Input Nodes before processing:")
    log.Println(nodeInputList)


    // Loop through the input nodes
    for j := range nodeInputList.Nodes {
       
        nodeName = nodeInputList.Nodes[j].Id
        //log.Printf(nodeName + " processing" )
 
        // Find the node properties from master node-properties file
        for i := range nodeList.Nodes {

            //log.Printf(nodeName + " finding properties" )

            if nodeList.Nodes[i].Id == nodeName {
                metricSQL = nodeList.Nodes[i].MetricSql
                influxURL = nodeList.Nodes[i].InfluxURL
                baseLine = nodeList.Nodes[i].Baseline
                //defaultThreshold = nodeList.Nodes[i].DefaultThreshold
              
                //log.Printf(nodeName + " found" )
                //log.Printf(metricSQL)
                //log.Printf(influxURL)
                //log.Printf("Baseline %d", baseLine)
                //log.Printf("DefaultThreshold %d", defaultThreshold)
     
                // set the baseline of the input node to baseline from master node properties
                //nodeInputList.Nodes[i].Baseline = strconv.Itoa(baseLine)
                nodeInputList.Nodes[i].Baseline = baseLine

                break
            }
        }

	q := client.Query{
		Command:  metricSQL,
		Database: "metrics",
	}

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{
		//Addr: "http://localhost:8086",
		//Addr: "https://metrics-shared.prod.company.com/",
		//Addr: "https://metricsengine-shared.prod.company.com/",
                Addr: influxURL,
	})
        if err != nil {
           log.Println("Influx DB Client error")
           c.AbortWithStatus(http.StatusNotFound)
        }

	if response, err := dbclient.Query(q); err == nil {
		if response.Error() != nil {
                    log.Println("Error in query results")
                    log.Println(response)
                    c.AbortWithStatus(http.StatusNotFound)
		}
		res := response.Results
                log.Println("got query results")
                log.Println(res)
		//fmt.Printf ("Datatype of result is: %T\n", res)
                //GetValueFromResult(res.([]byte)) // this does not work like this

		log.Println(res[0])
                fmt.Printf("res[0] Datatype %T \n", res[0])
                fmt.Printf("res[0].Series \n", res[0].Series)
                fmt.Println("---------------------\n\n")

//-------

                seriesData := res[0]

                for j := range seriesData.Series {
		    log.Println("Processing Series item index:")
		    log.Println(j)
		
		    for k := range seriesData.Series[j].Tags {

                       fmt.Println(seriesData.Series[j].Tags[k]) 
		    }
                    
                    for i := range seriesData.Series[j].Values {
		       fmt.Println("Value0:")
		       fmt.Println(seriesData.Series[j].Values[i][0])
		       fmt.Println("Value1:")
		       fmt.Println(seriesData.Series[j].Values[i][1])
		    }
		}

//-------

                // TEMP SETTING
                nodeInputList.Nodes[j].Value = -10

                // print the result to screen (to debug)
		c.JSON(http.StatusOK, res)
	} else {
                log.Println("failed to get query results")
                log.Println(response)
                c.AbortWithStatus(http.StatusNotFound)
	}
     
    //c.JSON(http.StatusOK,gin.H{ "message": "in GetNodeMetric!", } )

   }

   // print the input node list with updated values
   log.Println(nodeInputList)
  
   // write the updated input nodes json to a file again 
   file, _ := json.MarshalIndent(nodeInputList, "", " ")
   _ = ioutil.WriteFile("test.json", file, 0644)

}

func QueryDbLocalMac(c *gin.Context) {
        //cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"
        cmd := "SELECT sum(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"
        log.Println("Query Is", cmd)

	q := client.Query{
		Command:  cmd,
		Database: "metrics",
	}

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
        if err != nil {
           c.AbortWithStatus(http.StatusNotFound)
        }

	if response, err := dbclient.Query(q); err == nil {
		if response.Error() != nil {
                    log.Println("failed to get query results")
                    c.AbortWithStatus(http.StatusNotFound)
		}
		res := response.Results
                log.Println("got query results")
                log.Println(res)
		c.JSON(http.StatusOK, res)
	} else {
                c.AbortWithStatus(http.StatusNotFound)
	}
}

/*
func QueryDbNotWorking() (res OracleMetric) {
        cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"

	q := client.Query{
		Command:  cmd,
		Database: "metrics",
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			//return res, response.Error()
                        //return "message": "Query Error!"
		}
		res := response.Results
	} else {
		//return res, err
		return res
	}
	//return res, nil
	return res
}
*/

/*
func QueryDbOrig() (res []client.Result, err error) {
        cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"

	q := client.Query{
		Command:  cmd,
		Database: "metrics",
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
*/
