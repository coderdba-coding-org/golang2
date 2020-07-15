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

type TodosSlice []Todo

type TodoTagged struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type TodosTaggedSlice []TodoTagged

var todoList TodosSlice
var todoListTagged TodosTaggedSlice

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
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/printrequest", homeLinkPrintRequest).Methods("GET")
	router.HandleFunc("/todos", TodoIndex).Methods("GET")
	router.HandleFunc("/todostagged", TodoIndexTagged).Methods("GET")

	router.HandleFunc("/todos", PostTodoIndex).Methods("POST")
	router.HandleFunc("/todostagged", PostTodoIndexTagged).Methods("POST")
	// start 'server' web server
	log.Fatal(http.ListenAndServe(serverport, router))
}

func startClient() {
	// create routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", clientHomeLink).Methods("GET")
	router.HandleFunc("/todos", clientGetTodosFromServer).Methods("GET")
	router.HandleFunc("/todostagged", clientGetTodosFromServer).Methods("GET")
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
// not fully json like - the struct keys in the struct start with uppercase (no json tags in struct)
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := TodosSlice{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	oneMoreTodo := Todo{Name: "Sleep well"}

	todos = append(todos, oneMoreTodo)

	// write the todo struct to response as json
	json.NewEncoder(w).Encode(todos)
}

// my own function
// request body is a json todo item - one item per POST - which should be saved into a global variable
func PostTodoIndex(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println("PostTodoIndex: Failed to decode json request")
	}
	todoList = append(todoList, todo)
	fmt.Println(todoList)
}

func PostTodoIndexTagged(w http.ResponseWriter, r *http.Request) {
	var todo TodoTagged
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println("PostTodoIndexTagged: Failed to decode json request")
		fmt.Println(err)
	}

	fmt.Println(todo)

	todoListTagged = append(todoListTagged, todo)
	fmt.Println(todoListTagged)
}

// https://thenewstack.io/make-a-restful-json-api-go/
// fully json like - the real 'json tags' here start with lowercase
func TodoIndexTagged(w http.ResponseWriter, r *http.Request) {

	todos := TodosTaggedSlice{
		TodoTagged{Name: "Write presentation"},
		TodoTagged{Name: "Host meetup"},
	}
	// write the todo struct to response as json
	json.NewEncoder(w).Encode(todos)
}

func clientGetTodosFromServer(w http.ResponseWriter, r *http.Request) {

	// find the incoming URL and accordingly choose whether to get /todos or /todostagged from server
	// note: the URL will be just the endpoint like /todo and not the whole http://host:port/endpoint
	endpoint := r.URL.String()
	fmt.Println("clientGetTodosFromServer: URL is " + endpoint)

	// create a request to the server
	//req, err := http.NewRequest("GET", serverURL + "/todos", nil)
	req, err := http.NewRequest("GET", serverURL+endpoint, nil) // assuming same endpoints are used in client and server
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
