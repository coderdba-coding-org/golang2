package main

import (
        "github.com/gin-gonic/gin"
	//"errors"
	"log"
        //"encoding/json"
	"time"
        "net/http"

        "github.com/influxdata/influxdb/client/v2"
)

/*
type OracleMetric struct {
	time        time.Time 'json:TimeStamp'
	DbName      int 'json:DbName'
	MetricValue int 'json:MetricValue'
}
*/

type OracleMetric struct {
	time        time.Time 
	DbName      int 
	MetricValue int 
}

func main() {

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

        // db query endpoint
	r.GET("/query", QueryDb)
	//r.GET("/query", gin.H(queryDb))
	//r.GET("/query", func(c *gin.Context) { queryDb })


        // original port 8080
	err := r.Run("0.0.0.0:8080")
	//err := r.Run("0.0.0.0:5050")

	if err != nil {
	}

}

func QueryDb(c *gin.Context) {
        cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time = now() - 30m"
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
        cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time = now() - 30m"

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
        cmd := "SELECT mean(db_up_status) FROM gowrishora1 WHERE oracledb_service = 'ORDPROD' AND time = now() - 30m"

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
