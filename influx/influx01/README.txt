-- https://medium.com/spankie/getting-started-with-influxdb-with-golang-example-10990c5efee7

-- https://github.com/golang-standards/project-layout
-- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

-- Gin for API (with router): https://levelup.gitconnected.com/build-your-first-rest-api-in-go-language-using-gin-framework-827aadc14e07
-- Gin: https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
-- Gin with mySql - with positional parameters: https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8
-- Gin with Mongo: https://dev.to/faruq2/how-to-build-a-crud-rest-api-with-go-gin-and-fauna-37o6
-- Gin with client and server: https://medium.com/wesionary-team/building-rest-api-in-gin-framework-8c069716113e
-- Gin JSON binding using 'BindJSON' which can give EOF error: https://github.com/gin-gonic/gin/issues/715
-- Gin JSON binding using 'ShouldBind': https://github.com/gin-gonic/gin/issues/715

-- Print format cheat sheet: https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/

-- Influx client for go: https://golang.hotexamples.com/examples/github.com.influxdata.influxdb.influxql/-/ParseStatement/golang-parsestatement-function-examples.html

-- Fetch - Calling API from within react:  https://pusher.com/tutorials/consume-restful-api-react
-- Fetch with 'await': https://dev.to/johnpaulada/synchronous-fetch-with-asyncawait

-- CORS setting in Gin: https://ramezanpour.net/post/2020/08/23/cors-support-go-gin
-- CORS explained: https://www.pluralsight.com/guides/allow-access-control-origin-in-create-react-app
-- CORS explained: https://stackoverflow.com/questions/45975135/access-control-origin-header-error-using-axios-in-react-web-throwing-error-in-ch
-- CORS doc: https://github.com/gin-contrib/cors
-- CORS - cors package usage: https://github.com/gin-contrib/cors
-- CORS - cors package usage: https://skarlso.github.io/2016/02/02/doing-cors-in-go-with-gin-and-json/
-- CORS - another cors package: cors "github.com/rs/cors/wrapper/gin" --> https://stackoverflow.com/questions/64112975/access-to-xmlhttprequest-at-address-from-origin-address-has-been-blocked-by-cors/64113377
-- CORS - check if header is coming or not: https://aws.amazon.com/premiumsupport/knowledge-center/no-access-control-allow-origin-error/

-- Struct traversing to find key: https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
-- Struct to file as json: https://www.golangprograms.com/golang-writing-struct-to-json-file.html
-- Access JSON elements in go: https://stackoverflow.com/questions/35660467/how-to-access-fields-of-a-json-in-go
-- Access JSON elements in go: https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
-- Access JSON elements in go: https://yourbasic.org/golang/json-example/#decode-unmarshal-json-to-struct

-- Convert a number to json.Number:  https://stackoverflow.com/questions/46702871/how-to-convert-float64-to-json-number-golang
-- json.Number:  https://golang.org/pkg/encoding/json/#Number

-- Influx 'Result' documentation: https://pkg.go.dev/github.com/influxdata/influxdb/client/v2#Result
---- https://pkg.go.dev/github.com/influxdata/influxdb@v1.8.4/models#Row
---- https://pkg.go.dev/github.com/influxdata/influxdb@v1.8.4/models

-- Reading JSON file to Struct: https://www.golangprograms.com/golang-read-json-file-into-struct.html#:~:text=The%20json%20package%20includes%20Unmarshal,should%20be%20in%20capitalize%20format.
-- Reading JSON file to Struct: https://tutorialedge.net/golang/parsing-json-with-golang/
-- Maps, arrays: https://www.golang-book.com/books/intro/6

-- json.Number to float64 conversion:  https://stackoverflow.com/questions/48443495/convert-json-number-into-int-int64-float64-in-golang
	-- For error "interface conversion: interface {} is json.Number, not float64"

Initialize the project module
go mod init influx01

Next Create a Database. In your terminal type:
step1: influx // to enter influx DB mode
step2: create database go_influx // to create the db

Get the imports
go get net/http
go get github.com/go-chi/chi // http router
go get github.com/go-chi/chi/middleware // package for commonly used middlewares
go get github.com/sirupsen/logrus // cool logger package
go get github.com/spankie/go-influx/handlers // handler packages


