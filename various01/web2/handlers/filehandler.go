package handlers

import (
        "fmt"
        //"io/ioutil"
        "net/http"
        //"github.com/gorilla/mux"
        "os"
)


func createFile() (err error) {

    f, err := os.Create("/tmp/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return err
    }

    fmt.Println(l, "async bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}

func AsyncCreateFile(w http.ResponseWriter, r *http.Request) {

    go createFile()

}

func SyncCreateFile(w http.ResponseWriter, r *http.Request) {

    f, err := os.Create("/tmp/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
   

    fmt.Println(l, "bytes written successfully")
    fmt.Fprintf(w, "wrote file successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

}
