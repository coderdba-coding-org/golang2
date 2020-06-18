package main
  
import (
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
        "web1/pkg/db"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome home!")
}

func main() {

        //db.InitEvents()
        router := mux.NewRouter().StrictSlash(true)
        router.HandleFunc("/", homeLink)
        router.HandleFunc("/event", db.CreateEvent).Methods("POST")
        router.HandleFunc("/events", db.GetAllEvents).Methods("GET")
        router.HandleFunc("/events/{id}", db.GetOneEvent).Methods("GET")
        router.HandleFunc("/events/{id}", db.UpdateEvent).Methods("PATCH")
        router.HandleFunc("/events/{id}", db.DeleteEvent).Methods("DELETE")
        //router.HandleFunc("/events/{id}", db.deleteEvent).Methods("DELETE") // to check if lowercase function is accessible
        log.Fatal(http.ListenAndServe(":8081", router))

}
