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

        client "github.com/influxdata/influxdb/client/v2"
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

var nodeList NodeList

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
    
        // we initialize our Users array
        //var nodeList NodeList
    
        // we unmarshal our byteArray which contains our
        // jsonFile's content into 'NodesList' which we defined above
        json.Unmarshal(byteValue, &nodeList)
    
        // we iterate through every Node within our NodeList array and
        // print out the user Type, their name, and their facebook url
        // as just an example
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


        // db query endpoint
	//r.GET("/query", QueryDbLocalMac)
	r.GET("/query", QueryDbTgt)

        // does not work
	//r.GET("/query", gin.H(queryDb))
	//r.GET("/query", func(c *gin.Context) { queryDb })

        // get the metric value for a node
	//r.GET("/getmetric/:nodeName", GetNodeMetric)

        // original port 8080
	//err := r.Run("0.0.0.0:8080")
	err = r.Run("0.0.0.0:8080")
	//err := r.Run("0.0.0.0:5050")

	if err != nil {
	}

}

func QueryDbTgt(c *gin.Context) {
        //cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time > now() - 30m"
        cmd := "SELECT count(metric_value) FROM oracle_dbstatus WHERE _blossom_id = 'CI1234' AND application = 'RDBMS' AND oracledb_service = 'ORDPROD' AND metric_name = 'instance_status' AND time > now() - 30m"
        log.Println("Query Is", cmd)

	q := client.Query{
		Command:  cmd,
		Database: "metrics",
	}

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		//Addr: "https://metrics-shared.prod.company.com/",
              
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
}

func GetNodeMetric(nodeName string) {
     
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
