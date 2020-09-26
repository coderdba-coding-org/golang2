package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "gopkg.in/goracle.v2"
)

type DbConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

type DbName struct {
        DbName string
}

func getRows(db *sql.DB, dbQuery string) {

	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

        //fmt.Println(rows) // this prints bytes

	fmt.Println("... Parsing query results")

	var rowString string
	for rows.Next() {
		rows.Scan(&rowString)
		fmt.Println(rowString)
	}

}

func main() {

	//------------------------------
	// credentials hardcoded (original code style)
	//username := "TEMP"
	//password := "WRONGPW"
	//host := "db22-scan1:1521"
	//database := "TESTDB"
	//------------------------------

	//------------------------------
	// credentials from config file
	configfile := "config/oradb.json"
	dbConfig := DbConfig{}

	// TBD - Error handling
	file, err := os.Open(configfile)
	//if err != nil {  return err }
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbConfig)
	//if err != nil {  return err }

	fmt.Println("printing config read from file")
	fmt.Println(dbConfig)

	username := dbConfig.Username
	password := dbConfig.Password
	host := dbConfig.Host
	database := dbConfig.Database

	//------------------------------

	currentTime := time.Now()
	fmt.Println("Starting at : ", currentTime.Format("03:04:05:06 PM"))

	fmt.Println("... Setting up Database Connection")
	db, err := sql.Open("goracle", username+"/"+password+"@"+host+"/"+database)
	if err != nil {
		fmt.Println("... DB Setup Failed")
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("... Opening Database Connection")
	if err = db.Ping(); err != nil {
		fmt.Println("Error connecting to the database: %s\n", err)
		return
	}
	fmt.Println("... Connected to Database")

	//------------------------------
        // Querying a single column works
	//------------------------------
	//dbQuery := "select table_name from user_tables where table_name not like 'DM$%' and table_name not like 'ODMR$%'"
	dbQuery := "select name from v$database"
        //dbQuery := "select b.machine from v$access a, v$session b where a.sid=b.sid order by b.machine"

        // multi column query not working
        //dbQuery := "select b.machine, a.owner, a.type, a.object, b.sid, b.username from v$access a, v$session b where a.sid=b.sid order by machine, username, type, owner, object"

	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("... Parsing query results")
	var rowString string
	for rows.Next() {
		rows.Scan(&rowString)
		fmt.Println(rowString)
	}

	//----------------------------------
        // Querying as json - single column
	//----------------------------------
        dbQuery = "select json_object('DbName' is name) database_name from v$database"
        //dbQuery = "select json_object(name) from v$database"

	rows, err = db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

        //fmt.Println(rows) // this prints bytes

	fmt.Println("... Parsing query results")
	//var rowString string
	for rows.Next() {
		rows.Scan(&rowString)
		fmt.Println(rowString)
	}

	//--------------------------------------------------
        // Querying as json - single column, single row
	//--------------------------------------------------
        dbQuery = "select json_object('DbName' is name) database_name from v$database"

	rows, err = db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

        //fmt.Println(rows) // this prints bytes

	fmt.Println("... Parsing query results")
	//var rowString string
	for rows.Next() {
		rows.Scan(&rowString)
		fmt.Println(rowString)
	}

	//------------------------------------------------
        // Querying as json - single column, multi row
	//------------------------------------------------
        dbQuery = "select b.machine from v$access a, v$session b where a.sid=b.sid order by b.machine"
        getRows(db, dbQuery)

	//------------------------------
        // Close and exit
	//------------------------------
	fmt.Println("... Closing connection")
	finishTime := time.Now()
	fmt.Println("Finished at ", finishTime.Format("03:04:05:06 PM"))

}

