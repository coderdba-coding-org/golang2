package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/formdata", printFormData).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// Postman Body: x-www-form-urlencoded
func printFormData(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Println("Username: ", r.FormValue("username"))
	fmt.Println("Password: ", r.FormValue("password"))

	type responseStruct struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	responseValue := responseStruct{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	//w.Write(responseValue)
	json.NewEncoder(w).Encode(responseValue)

}
