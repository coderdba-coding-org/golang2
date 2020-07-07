### Source
https://www.sohamkamani.com/golang/2018-06-24-oauth-with-golang/   

https://github.com/sohamkamani/go-oauth-example  

http://networkbit.ch/golang-http-client/

Main reference: https://www.sohamkamani.com/golang/2018-06-24-oauth-with-golang/

For extended http code: http://networkbit.ch/golang-http-client/

For printing http requests to debug: https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000

To place token in header: https://developer.github.com/v3/ 

Response to byte-array: https://stackoverflow.com/questions/39945968most-efficient-way-to-convert-io-readcloser-to-byte-array

Response to byte-array: https://stackoverflow.com/questions/41815909/is-there-a-way-to-extract-json-from-an-http-response-without-having-to-build-str


### Setup oauth application in Github
Settings --> Developer Settings --> OAuth Apps

There, provide this:

Application name: app1 (or other)

Homepage URL: http://localhost:8080

Authorization callback URL: http://localhost:8080/oauth/redirect (or other endpoint of the app which will receive the access_token returned by Github)  


### Code Notes
Replace these in main.go:  

const clientID = "<>"  

const clientSecret = "<>"  

### Flow
First login page appears

Click and it will take to github

There, logon to github

Then github redirects as follows: http://localhost:8080/oauth/redirect?code=e90ef6b262ca74456fcb

The 'code' in the above redirection is the 'request token' - to process the upcoming (see below) request to get 'access token'

Then in main.go, "/oauth/redirect" handler gets the 'code' part

and sends another request to github with this code, client-id and client-secret: 

("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)

The return value from Github will be the 'access token'


