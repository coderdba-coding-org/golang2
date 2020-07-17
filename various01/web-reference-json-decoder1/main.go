package main

// https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func get_content() {
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	// READALL WAY OF JSON RESPONSE MARSHALLING
	// THIS IS WORKING
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(" Error reading response body. ", err)
	}
	fmt.Printf("clientGetTodosFromServer: Response Body\n %s\n", body)

	// DECODER WAY OF JSON RESPONSE MARSHALLING
	// THIS IS NOT WORKING

	decoder := json.NewDecoder(res.Body)

	var data Tracks

	err = decoder.Decode(&data)
	if err != nil {
		//fmt.Printf("Error decoding json body")

		//fmt.Printf("%T\n%s\n%#v\n",err, err, err)
		//switch v := err.(type){
		//case *json.SyntaxError:
		//fmt.Println(string(body[v.Offset-40:v.Offset]))
		//}
	}

	fmt.Println(data)

	for i, track := range data.Toptracks.Track {
		fmt.Printf("%d: %s %s\n", i, track.Artist.Name, track.Name)
	}
}

func main() {
	get_content()
}
