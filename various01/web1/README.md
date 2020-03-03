## REFERENCES
Build a RESTful JSON API with GOlang  
- https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da  

## NOTES
Many functions - not defined in main.go - have been named with Uppercase starting letter - so that they can be accessed from the main program  

## STEPS
$ go mod init web1  
$ go get -u github.com/gorilla/mux  

Then code the rest of the stuff  

$ go build  
$ web1 (which is the executable)  

In web browser - http://localhost:8082 (or other port used)  
