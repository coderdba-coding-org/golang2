package main

import (
"io"
"log"
"net/http"
)

func main(){

// create a handler
helloHandler := func(w http.ResponseWriter, r *http.Request) {
  log.Print("endpoing called")
  io.WriteString(w, "hello world\n")
}

// registr the handler
http.HandleFunc("/hello", helloHandler)

// start the server
log.Fatal(http.ListenAndServe(":8888", nil))

}

