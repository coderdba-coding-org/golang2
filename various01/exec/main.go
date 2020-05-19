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
        "os/exec"
        //"path"
        "bytes"
        //"strings"
        //"time"
)

func main() {

        log.Println("In main()")

        runTestShellCommand1()
        runTestShellCommand2()

        // Maintenance loop
        //for {
                //log.Println("Loop check and fix")
                ////fmt.Println("Loop check and fix")

                //// Maintenance functions
                //runTestKubectl1()
                //patroniReinit()

                //time.Sleep(300 * time.Second)
        //}
}


func runTestShellCommand1() {
        out, err := exec.Command("date").Output()
        if err != nil {
                log.Fatal(err)
        }
        // Printf prints string - println prints bytes
        fmt.Println("===================")
        fmt.Printf("The date fmt.Printf is %s\n", out)
        fmt.Println("The date fmt.Println is %s\n", out) //println does not recognize formatter %s

        fmt.Println("===================")
        fmt.Printf("The date string(out) fmt.Printf is %s ", string(out))
        fmt.Println("The date string(out) fmt.Println is %s\n", string(out)) //println does not recognize formatter %s

        fmt.Println("===================")
        log.Printf("The date log.Printf is %s\n", out)
        log.Println("The date log.Println is %s\n", out) //println does not recognize formatter %s

        log.Println("===================")
        log.Printf("The date string(out) log.Printf is %s ", string(out))
        log.Println("The date string(out) log.Println is %s\n", string(out)) //println does not recognize formatter %s

        log.Println("===================")
}

func runTestShellCommand2() {
        //shellcmd := "ls"
        shellcmd := "date"

        cmd := exec.Command("bash", "-c", shellcmd)
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Command output out.String() Printf: %q\n", out.String())
        fmt.Println("Command output  out.String() Println: %q\n", out.String())

}

func runTestKubectl1() {

        log.Println("In runTestKubectl1()")

        cmd := fmt.Sprintf("kubectl get nodes")

        fmt.Printf("Command is %s\n", cmd)
        log.Print("Command is %s\n", cmd)

        out, err := exec.Command("bash", "-c", cmd).Output()

        if err != nil {
                log.Println(err)
                //log.Fatal(err)
        }

        fmt.Printf("Output is %s", out)
        log.Printf("Output is %s\n", out)

}
