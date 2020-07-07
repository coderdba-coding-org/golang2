package main

// Main reference: https://www.sohamkamani.com/golang/2018-06-24-oauth-with-golang/
// For extended http code: http://networkbit.ch/golang-http-client/
// For printing http requests to debug: https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//const clientID = "<your client id>"
//const clientSecret = "<your client secret>"

const clientID = "9d1a3826c6b575840df3"
const clientSecret = "2697934ff00d15da664948c41f2e3dbb1f5566b4"

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type UserEndpointResponse struct {
	Login    string `json:"login"`
	Location string `json:"location"`
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// We will be using `httpClient` to make external HTTP requests later in our code
	httpClient := http.Client{}

	// Create a new redirect route route
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		// First, we need to get the value of the `code` query param
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		// Next, lets for the HTTP request to call the github oauth enpoint
		// to get our access token
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		// We set this header since we want the response as JSON
		req.Header.Set("accept", "application/json")

		// Send out the HTTP request
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// SOMEHOW, PRINTING THIS IS CAUSING THE BODY TO LOSE SOMETHING
		// AND THE NEXT NewDecoder stuff is not seeing the token in the body
		// print the response body
		//resbody, err := ioutil.ReadAll(res.Body)
		//if err != nil {
		//	log.Fatal("Error reading body. ", err)
		//}
		//fmt.Printf("%s\n", resbody)

		// Parse the request body into the `OAuthAccessResponse` struct -
		var t OAuthAccessResponse

		// see "&t" below
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Println("access token is: " + t.AccessToken)
		fmt.Fprintf(os.Stderr, "access token is: "+t.AccessToken)

		// GET USER DATA USING THE AUTH TOKEN (you can do other API calls also to github, but doing user here)
		// new code (not in the original example)
		// From http://networkbit.ch/golang-http-client/
		// and referring to https://developer.github.com/v3/

		// create request to 'user' endpoint of github api
		req, err = http.NewRequest("GET", "https://api.github.com/user", nil)

		// We set this header since we want the response as JSON
		req.Header.Add("accept", "application/json")
		if err != nil {
			log.Fatal("Error creating request. ", err)
		}

		// add the access-token to header
		req.Header.Add("Authorization", "token "+t.AccessToken)
		if err != nil {
			log.Fatal("Error creating request. ", err)
		}

		// debug - print the request details
		fmt.Printf("--> %s\n\n", formatRequest(req))

		// create a http client with timeout
		client := &http.Client{Timeout: time.Second * 10}

		// run the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading body. ", err)
		}

		// print the response body
		fmt.Printf("%s\n", body)

		// THIS TYPE OF DECODING AND PRINTING NEEDS THE STRUCT TO HAVE ALL FIELDS IN THE JSON
		// print only the 'login' (which is the userid of the user) from the response body
		//var u UserEndpointResponse
		//if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		//	fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//fmt.Println("login is: " + u.Login)
		//fmt.Println("location is: " + u.Location)

		// THIS IS PART OF ORIGINAL CODE from https://www.sohamkamani.com/golang/2018-06-24-oauth-with-golang/
		// Finally, send a response to redirect the user to the "welcome" page
		// with the access token
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
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
