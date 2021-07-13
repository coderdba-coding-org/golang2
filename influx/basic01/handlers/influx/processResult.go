// Process influx result json obtained by querying influx

package influx

import (
	"log"
	"fmt"
	"encoding/json"
	//"strings"
	//"errors"

	client "github.com/influxdata/influxdb1-client/v2"
)


// To process query results like the following:
/*
* Query:  "select value from cpu" --> which returns just the metric column 'value' in 'cpu' table (one or more rows outputted)
* Resultset printed in browser:
*   [{"statement_id":0,"Series":[{"name":"cpu","columns":["time","value"],"values":[["2020-02-26T10:51:04.788081Z",0.64],["2020-02-26T10:51:10.30838Z",0.64]]}],"Messages":null}]
* From golang log/println: 
*   2021/07/13 16:07:53 [{0 [{cpu map[] [time value] [[2020-02-26T10:51:04.788081Z 0.64] [2020-02-26T10:51:10.30838Z 0.64]] false}] [] }]
*
* Logic:
*  For every item in the "Series" of the first row of the JSON (acutually Series is just one row altogether with many items in it)
*  This Series is the resultset's 0th element (res[0]) 
*   - theoretically this Series may have more than one element, but in this case only one element exists
*  For each element in the 'Series' 
*   - get 'Tags' - which may or may not be present (this one is not that important)
*   - get the 'Values' element which is an array or arrays.  Each inner array has 2 elements. Print those two elements
*
*/
func ProcessResult1(res []client.Result) {

            fmt.Println("ProcessResult1(): In ProcessResult1")

            // The first item in the resultset is the required data - and as json would have a tag 'Series'
            seriesData := res[0]

            // For each element of the 'Series' loop through the elements
            for k := range seriesData.Series {
                log.Println("Processing Series item index:")
                log.Println(k)

                // Get the tag of the element
                for l := range seriesData.Series[k].Tags {

                   // There may or may not be tags
                   fmt.Println("TAG: " + seriesData.Series[k].Tags[l])
                }

                // Get the 'Values' element and process the inner arrays
                // For each inner array of 'Values' array
                for m := range seriesData.Series[k].Values {
               
                   fmt.Printf("**** Processing row/record: %d \n",  m)
                   fmt.Println("Value0:")
                   fmt.Println(seriesData.Series[k].Values[m][0])
                   fmt.Println("Value1:")
                   fmt.Println(seriesData.Series[k].Values[m][1])

                   // convert the numeric element (if any) to float from json.Number
                   var valueGotten float64
                   valueGotten, _ = seriesData.Series[k].Values[m][1].(json.Number).Float64()
                   fmt.Println("valueGotten after float conversion is:")
                   fmt.Println(valueGotten)
                }
            }
}
