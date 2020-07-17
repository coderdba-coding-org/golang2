package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//import "io/ioutil"

type Tracks struct {
	Toptracks Toptracks_info
}

type Toptracks_info struct {
	Track []Track_info
	Attr  Attr_info `json: "@attr"`
}

type Track_info struct {
	Name       string
	Duration   string
	Listeners  string
	Mbid       string
	Url        string
	Streamable Streamable_info
	Artist     Artist_info
	Attr       Track_attr_info `json: "@attr"`
}

type Attr_info struct {
	Country    string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}

type Streamable_info struct {
	Text      string `json: "#text"`
	Fulltrack string
}

type Artist_info struct {
	Name string
	Mbid string
	Url  string
}

type Track_attr_info struct {
	Rank string
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

const serverport = ":8080"

//const clientport = ":8081"
//const serverURL = "http://localhost" + serverport
//const clientURL = "http://localhost" + clientport

func startserver() {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/getlocal", getLocal).Methods("GET")
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/getweb1", getWebReadAll).Methods("GET")
	//router.HandleFunc("/getweb2", getWebDecoder).Methods("GET")
	log.Fatal(http.ListenAndServe(serverport, router))
}

func get_content() {
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data Tracks
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		//switch v := err.(type){
		//case *json.SyntaxError:
		//fmt.Println(string(body[v.Offset-40:v.Offset]))
		//}
	}

	fmt.Println("Json decoded successfully")
	fmt.Println("Printing data in the json")
	for i, track := range data.Toptracks.Track {
		fmt.Printf("%d: %s %s\n", i, track.Artist.Name, track.Name)
	}
}

func get_content_2() error {
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("get_content_2: Error reading response body. ", err)
	}

	// print the response body - to console (not as response out)
	fmt.Printf("clientGetTodosFromServer: Response Body\n %s\n", body)

	// Somehow this decoder does not seem to show any data
	fmt.Println("\n\nUsing json.NewDecoder now")
	fmt.Println("==============================")
	decoder := json.NewDecoder(res.Body)
	var data Tracks
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Error decoding response body:  %T\n%s\n%#v\n", err, err, err)

		switch v := err.(type) {
		case *json.SyntaxError:
			fmt.Println(string(body[v.Offset-40 : v.Offset]))
		default:
			fmt.Printf("Error type is: %v", v)
		}

		return errors.New("Error decoding response body")
	}

	fmt.Println("Json decoded successfully")
	fmt.Println("Printing data in the json")
	for i, track := range data.Toptracks.Track {
		fmt.Printf("%d: %s %s\n", i, track.Artist.Name, track.Name)
	}

	return nil
}

// web functions

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getWebReadAll(w http.ResponseWriter, r *http.Request) {

	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("get_content_2: Error reading response body. ", err)
	}

	// print the response body - to console (not as response out)
	//fmt.Printf("getWeb: Response Body\n %s\n", body)

	// send back response using this body
	w.Write(body)
}

func main() {
	//get_content()

	//get_content_2()

	startserver()
}
