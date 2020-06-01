## REFERENCES

This has these:  
1. Using go routine to start a background job  
2. Using global variable to state the status of background job/goroutine "functionality" (not whether that job/goroutine is up or not thing)  
3. TBD - to check whether the job/goroutine is even up or not  

Build a RESTful JSON API with GOlang  
- https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da  

For file handling:  
https://golangbot.com/write-files/  

For goroutines:  
This or some other http://codesolid.com/post/optimizing-aws-s3-uploads-with-golang-and-goroutines/   

## NOTES  
Many functions - not defined in main.go - have been named with Uppercase starting letter - so that they can be accessed from the main program    

## STEPS
$ go mod init web4  
$ go get -u github.com/gorilla/mux  

Then code the rest of the stuff  

### To just run 
go run .

### To build and run
$ go build  
$ web4 (which is the executable)  

In web browser - http://localhost:8081 (or other port used)  
