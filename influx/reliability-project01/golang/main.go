package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	// local imports
	influx "dbtrends/handlers/influx"
	reliability "dbtrends/handlers/reliability"
)

var defaultInfluxDB string = "metrics"

var defaultInfluxURL string = "https://localhost:8086/"

var sampleInfluxQuery1 string = "SELECT count(metric_value) AS value FROM oracle_dbstatus WHERE config_item = 'CI123' AND application = 'RDBMS' AND oracledb_service = 'ORDPROD' AND metric_name = 'instance_status' AND time > now() - 30m"
var sampleInfluxQuery2 string = "SELECT count(Avg_Elap_Tm_by_Exec_ms) FROM ( SELECT sum(tot_elap_sec) / sum(execs) * 1000 AS Avg_Elap_Tm_by_Exec_ms FROM oracle_topsql_elpsd WHERE (config_item = 'CI00417360' AND db_service = 'ORDPROD_NODE1' AND (schema = 'ORDERUSER' OR schema = 'ORDERSUSER') AND environment = 'production') AND (sql_text_short =~ /^insert into order_snapshot/) AND time > now() - 60m GROUP BY time(10m) fill(null)) WHERE Avg_Elap_Tm_by_Exec_ms <= 25"

func main() {

	log.Println("Entering main()")

	log.Println("main(): Starting web server")
	StartWebServer()

}

// Start web server
func StartWebServer() {

	// ROUTER WITH CUSTOM SETTINGS 2 (with "github.com/rs/cors/wrapper/gin")
	router := gin.Default()
	router.Use(cors.AllowAll())

	// homepage
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome!",
		})
	})

	// health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy!",
		})
	})

	// trial endpoint - to try out things quickly
	router.GET("/trial", func(c *gin.Context) {
		influx.TrialInflux()
		c.JSON(200, gin.H{
			"message": "Trial Done!",
		})
	})

	// trial endpoint - to try out things quickly
	router.GET("/trialquery1", func(c *gin.Context) {
		//influx.QueryInflux(defaultInfluxURL, sampleInfluxQuery1, defaultInfluxDB)
		res, err := influx.QueryInflux(defaultInfluxURL, sampleInfluxQuery1, defaultInfluxDB)
		if err == nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(501, gin.H{
				"message": "Trial Query Errored in Influx Query!",
			})
		}
		/*
			c.JSON(200, gin.H{
				"message": "Trial Query Done!",
			})
		*/
	})

	router.GET("/trialquery2", func(c *gin.Context) {
		//influx.QueryInflux(defaultInfluxURL, sampleInfluxQuery2, defaultInfluxDB)
		res, err := influx.QueryInflux(defaultInfluxURL, sampleInfluxQuery2, defaultInfluxDB)
		if err == nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(501, gin.H{
				"message": "Trial Query Errored in Influx Query!",
			})
		}
		/*
			c.JSON(200, gin.H{
				"message": "Trial Query Done!",
			})
		*/
	})

	router.GET("/trialreliability1", func(c *gin.Context) {
		//reliability.CalculateReliability("sql_db")
		reliability.CalculateReliability("db_up")
		c.JSON(200, gin.H{
			"message": "Trial Reliability Done!",
		})
	})

	router.GET("/reliability/:component", func(c *gin.Context) {
		component := c.Param("component")
		reliability.CalculateReliability(component)
		c.JSON(200, gin.H{
			"message": "Reliability calculated for: " + component,
		})
	})

	router.GET("/nodereliability/:component", func(c *gin.Context) {
		component := c.Param("component")
		reliability.CalculateNodeReliability(component)
		c.JSON(200, gin.H{
			"message": "Node Reliability calculated for: " + component,
		})
	})

	router.GET("/nodelayerreliability/:component", func(c *gin.Context) {
		component := c.Param("component")
		reliability.CalculateNodeLayerReliability(component)
		c.JSON(200, gin.H{
			"message": "Node Layer Reliability calculated for: " + component,
		})
	})

	router.GET("/entityreliability/:component", func(c *gin.Context) {
		component := c.Param("component")
		reliability.CalculateEntityReliability(component)
		c.JSON(200, gin.H{
			"message": "Entity Reliability calculated for: " + component,
		})
	})

	router.Run(":8080")
	log.Println("StartWebServer(): Started web server")

}
