package main

import (
    "fmt"
    "time"
    "database/sql"
    _ "gopkg.in/goracle.v2"
)

func main(){
    username := "TEMP";
    password := "WRONGPW"
    host := "db22-scan1:1521";
    database := "TESTDB";

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
