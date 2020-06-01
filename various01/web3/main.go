package main

import (
	"fmt"
	"log"
	"net/http"
        "time"

	"github.com/gorilla/mux"

        //"web3/pkg/db"
        "web3/handlers"
)

func reconcileServiceLoop() {
   for {
        fmt.Println("Infinite Loop 1")
        time.Sleep(time.Second)
    }
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func startWebServer() {

        fmt.Println("Starting web router")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	router.HandleFunc("/createfilesync", handlers.SyncCreateFile).Methods("GET")
	router.HandleFunc("/createfileasync", handlers.AsyncCreateFile).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {

        // Call reconcileServiceLoop as a goroutine - and then call router directly (not via a function)
        // OR
        // Call startWebServer as goroutine first and then reconcileServiceLoop as non-goroutine function`

        //fmt.Println("Calling reconcileServiceLoop()")
	//reconcileServiceLoop()
        //go reconcileServiceLoop()

        //fmt.Println("Starting web router")
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	//router.HandleFunc("/createfilesync", handlers.SyncCreateFile).Methods("GET")
	//router.HandleFunc("/createfileasync", handlers.AsyncCreateFile).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8081", router))

        fmt.Println("Starting web router - via a function")
        go startWebServer()
        // this non-goroutine way will not let control go to the next steps
        //startWebServer()

        fmt.Println("Calling reconcileServiceLoop()")
	reconcileServiceLoop()
        // this goroutine way seems to not work with goroutine way of calling startWebServer
        //go reconcileServiceLoop()

}
