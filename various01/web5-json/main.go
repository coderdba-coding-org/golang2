package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	//"os"
	//"strings"
	//"time"
)

const serverport = ":8080"
const clientport = ":8081"
const serverURL = "http://localhost" + serverport
const clientURL = "http://localhost" + clientport

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

func main() {

	// start a file server to serve http files
	startFileServer()

	// start server http server
	go startServer()

	// start client http server
	startClient()
}

func startFileServer() {
	fs := http.FileServer(http.Dir("http"))
	http.Handle("/", fs)
}

func startServer() {
	// create routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/printrequest", homeLinkPrintRequest)
	router.HandleFunc("/todos", TodoIndex)
	// start 'server' web server
	log.Fatal(http.ListenAndServe(serverport, router))
}

func startClient() {
	// create routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", clientHomeLink)
	router.HandleFunc("/todos", clientGetTodosFromServer)

	// start 'client' web server
	log.Fatal(http.ListenAndServe(clientport, router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func clientHomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome client home!")
}

// https://thenewstack.io/make-a-restful-json-api-go/
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func clientGetTodosFromServer(w http.ResponseWriter, r *http.Request) {

	// create a request
	req, err := http.NewRequest("GET", serverURL+"/todos", nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "clientGetTodosFromServer: Error creating HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// header to request json response
	req.Header.Add("accept", "application/json")

	// send the request
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "clientGetTodosFromServer: Error from HTTP request sent: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// get the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("clientGetTodosFromServer: Error reading response body. ", err)
	}

	// print the response body - to console (not as response out)
	fmt.Printf("clientGetTodosFromServer: Response Body\n %s\n", body)

	// send the response
	w.Write(body)
}

func homeLinkPrintRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TBD this is TBD")
}

// formatRequest generates ascii representation of a request
// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
