package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	// local imports
	influx "basic01/handlers/influx"
)

//var defaultInfluxDB string = "metrics"
//var defaultInfluxURL string = "https://metricsengine.prod.company.com/"

var defaultInfluxDB string = "cpu1"
var defaultInfluxURL string = "http://localhost:8086"

var sampleInfluxQuery1 string = "SELECT * from cpu"
var sampleInfluxQuery2 string = "select value from cpu"

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

	router.Run(":8085")
	log.Println("StartWebServer(): Started web server")

}
