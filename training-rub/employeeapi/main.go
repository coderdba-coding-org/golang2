responsewriter - 
header.set(content-type, app-json)
writeheader(http.statusok)
json.NewEncoder(w).Encode(h.DB.GetEmployeeList)

type Employee struct {
FirstName string
LastName  string
EmpId     int
}

empSlice []Employee

type DB interface
{
GetFirstName(for an empid) returns string 
GetLastName(for an empid) returns string 
GetEmpId(for a name) returns int
GetEmployeeList(for all employees) returns empSlice
}

type PgDB struct {
dbConn *sql.DB --> This comes from import "database/sql"
}

The PostgresDB struct does not need any more fields as its main thing is connection
- it implements DB interface
- imports:
-- database/sql - for postgres db - this just exports the interface - and the commuity had write dribers for postgres
-- "github.com/lib/pq" - the library for postgres that uses database/sql interface

import (
_"github.com/lib/pq" --> The underscore before this is so that we dont call functions in it directly - instead use the functions defined for the interface database/sql
)

var pgdb PgDB
pgdb.DBSetup() // a method that references *PgDB or PgDB type

Connect to db, if err is not nil (panic(err))

rows, err := pg.dbconn.Query("select * from public.employee)

defer rows.Close() // we dont close the connection here - just the queried rows
for rows.Next() {
var  emp model.Employee // the struct
err = rows.Scan(&emp.EmpId, &emp.FirstName, &emp.LastName)
if err....

empSlice = append(empSlice, emp)
}
// check any error during this iteration over rows
err = rows.Err()
if err != nil {
}
return empSlice

=========================

API Client - a clinet to talk to apis

Reading response bnod - ioutil.Readall(resp.Body)
client := &http.Client{Transport: tr}
where 
tr := &http.Transport{
MaxIdleCons: 10
IdleConnTimeout: 30* time.Second
DisableCompression: true
}

===============

Project task:

