-- https://medium.com/spankie/getting-started-with-influxdb-with-golang-example-10990c5efee7

-- https://github.com/golang-standards/project-layout
-- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

-- Gin for API: https://levelup.gitconnected.com/build-your-first-rest-api-in-go-language-using-gin-framework-827aadc14e07
-- Gin: https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
-- Gin with mySql: https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8
-- Gin with Mongo: https://dev.to/faruq2/how-to-build-a-crud-rest-api-with-go-gin-and-fauna-37o6
-- Gin with client and servrer: https://medium.com/wesionary-team/building-rest-api-in-gin-framework-8c069716113e
-- Gin JSON binding using 'BindJSON' which can give EOF error: https://github.com/gin-gonic/gin/issues/715
-- Gin JSON binding using 'ShouldBind': https://github.com/gin-gonic/gin/issues/715

-- Reading JSON file to Struct: https://www.golangprograms.com/golang-read-json-file-into-struct.html#:~:text=The%20json%20package%20includes%20Unmarshal,should%20be%20in%20capitalize%20format.
-- Reading JSON file to Struct: https://tutorialedge.net/golang/parsing-json-with-golang/

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


