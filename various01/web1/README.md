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

## GET, POST ETC

### GET
localhost:8081/events  
localhost:8081/event/ID --> localhost:8081/event/1 and such  

### POST Using Postman
URL = localhost:8081/event  
Body type = JSON (application/json)  
Content = {"ID":"2","Title":"22","Description":"222"}  

### All events
router.HandleFunc("/", homeLink)  
router.HandleFunc("/event", db.CreateEvent).Methods("POST")  
router.HandleFunc("/events", db.GetAllEvents).Methods("GET")  
router.HandleFunc("/events/{id}", db.GetOneEvent).Methods("GET")  
router.HandleFunc("/events/{id}", db.UpdateEvent).Methods("PATCH")  
router.HandleFunc("/events/{id}", db.DeleteEvent).Methods("DELETE")  
