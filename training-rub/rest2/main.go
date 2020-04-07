-note: r is the http request 
switch r.Method {

case http.MethodGet:
	getEmployee()

case http.MethodPost:
	postEmployee()

}
