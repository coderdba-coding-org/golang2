package main

import (
        //"crypto/tls"
        //"encoding/base64"
        //"encoding/json"
        //"flag"
        "fmt"
        //"www.github.com/coderdba-coding-org/exec/app/logger"
        //"io/ioutil"
        "log"
        //"net"
        //"net/http"
        //"os"
        //"os/exec"
        //"path"
        //"bytes"
        "strings"
        //"time"
)

func main() {

        log.Println("In main()")

       
        // Multiline string
        // use backtick ` to define multiline string
        multilineString1 := `abc
def
ghi1`

	fmt.Printf(multilineString1)

        // now, split it into an array
        splitStrings := strings.Split("multilineString1", "\n")

        // this print errors out as this is []string - EXPECTED - so, OK
	fmt.Printf(splitStrings)
}

