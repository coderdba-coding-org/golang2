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
	configfile := "config/oratestdb.json"
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

}
