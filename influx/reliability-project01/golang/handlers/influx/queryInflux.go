// Query Influx DB and return result

package influx

import (
	"log"
	//"encoding/json"
	//"strings"
	"errors"

	client "github.com/influxdata/influxdb1-client/v2"
)

func TrialInflux() {
	log.Println("TrialInflux(): In Trial Influx")
}

func QueryInflux(url string, sql string, db string) (res []client.Result, err error) {

	log.Println("QueryInflux(): Entering the function")
	log.Println("QueryInflux(): URL: " + url)
	log.Println("QueryInflux(): sql: " + sql)
	log.Println("QueryInflux(): db: " + db)

	var blankResult []client.Result

	dbclient, err := client.NewHTTPClient(client.HTTPConfig{Addr: url})
	if err != nil {
		log.Println("Influx DB Client error")
		return blankResult, errors.New("QueryInflux(): Influx DB Client error")
	}

	q := client.Query{
		Command:  sql,
		Database: db,
	}

	response, err := dbclient.Query(q)
	if err == nil {
		if response.Error() != nil {
			log.Println("Error querying Influx")
			log.Println(response)
			return blankResult, errors.New("QueryInflux(): Error querying Influx")
		}
		res := response.Results

		//log.Println("QueryInflux(): -------------------------")
		log.Println("QueryInflux(): Got Influx query results:")
		log.Println(res)
		// Response datatype *client.Response, Response.Results datatype []client.Result
		log.Printf("QueryInflux(): Response datatype %T, Response.Results datatype %T", response, res)
		//log.Println("QueryInflux(): -------------------------")

		return res, nil

	} else {
		log.Println("QueryInflux(): Failed to get Influx query results")
		log.Println(response)
		return blankResult, errors.New("QueryInflux(): Failed to get Influx query results")
	}

	return blankResult, errors.New("QueryInflux(): Error querying Influx")

}
