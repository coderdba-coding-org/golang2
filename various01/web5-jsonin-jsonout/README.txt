References:
Basic web with http and router & JSON: https://thenewstack.io/make-a-restful-json-api-go/ (has serving a json)
Gorilla mux vs http for routes: https://levelup.gitconnected.com/experiment-golang-http-builtin-and-related-popular-packages-1d9a6dcb80d
Various rendering - json, xml etc: https://www.alexedwards.net/blog/golang-response-snippets#json

Web, Json with Struct+slice https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get

Slices: https://stackoverflow.com/questions/18042439/go-append-to-slice-in-struct


For extended http code: http://networkbit.ch/golang-http-client/
Parse to byte-array/map and struct: https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
Printing requests to debug: https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
Parsing json to: https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968

IOUtil to get response body in json: https://stackoverflow.com/questions/41815909/is-there-a-way-to-extract-json-from-an-http-response-without-having-to-build-str
IOUtil to get response body in json: https://stackoverflow.com/questions/39945968/most-efficient-way-to-convert-io-readcloser-to-byte-array

To place token in header: https://developer.github.com/v3/
Response to byte-array: https://stackoverflow.com/questions/39945968/most-efficient-way-to-convert-io-readcloser-to-byte-array
Response to byte-array: https://stackoverflow.com/questions/41815909/is-there-a-way-to-extract-json-from-an-http-response-without-having-to-build-str

What is map[string]interface{}
https://stackoverflow.com/questions/48988823/difference-between-mapstringinterface-and-interface  


===================
POST DATA TO SERVER
===================
URL: http://localhost:8080/todostagged
Method: POST
Body: raw --> JSON (application/json)
{key: value, key: value}
Example:
{"Name": "Write presentation"} --> only one key specified, other two left out
{"Name": "sleep well", "Completed": true} --> only two keys specified, one left out

{"Name": "clean desk", "Completed": false, "Due": "2018-03-29T13:34:00.000"} 
    --> DATE NOT WORKING

{"Name": "drive back", "Completed": false, "Due": "2018-03-29T13:34:26+05:30"} 
    --> {drive back false 2018-03-29 13:34:26 +0530 IST}

{"Name": "drive back", "Completed": false, "Due": "2018-03-29T13:34:26+05:00"}
    --> {drive back false 2018-03-29 13:34:26 +0500 +0500}

{"Name": "drive home", "Completed": false, "Due": "2018-03-29T13:34:00+00:00"} 
    --> date kind of works like this:
    {drive home false 2018-03-29 13:34:00 +0000 +0000}

{"Name": "drive back", "Completed": false, "Due": "2018-03-29T13:34:00+00:00"} 
    --> date kind of works like this:
    {drive back false 2018-03-29 13:34:00 +0000 +0000}

For UTC I guess this way:
{"Name": "drive back", "Completed": false, "Due": "2018-03-29T13:34:00Z"} 
    --> date kind of works like this: (this is incorrect usage)
    {drive back false 2018-03-29 13:34:00 +0000 UTC}


