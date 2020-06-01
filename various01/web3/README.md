## REFERENCES
This has:  
1. Goroutine to call a background job before starting web server  
--> useful for batch-job programs which also must have a web-endpoint like 'health'  

Build a RESTful JSON API with GOlang  
- https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da  

For file handling:  
https://golangbot.com/write-files/  

For goroutines:  
This or some other http://codesolid.com/post/optimizing-aws-s3-uploads-with-golang-and-goroutines/   

## NOTES  
Many functions - not defined in main.go - have been named with Uppercase starting letter - so that they can be accessed from the main program    

## STEPS
$ go mod init web3  
$ go get -u github.com/gorilla/mux  

Then code the rest of the stuff  

### To just run 
go run .

### To build and run
$ go build  
$ web3 (which is the executable)  

In web browser - http://localhost:8081 (or other port used)  
