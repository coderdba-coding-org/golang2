package main

import (
	//"crypto/tls"
	//"encoding/base64"
	//"encoding/json"
	//"flag"
	"fmt"
	//"io/ioutil"
	"log"
	//"net"
	//"net/http"
	//"os"
	"os/exec"
	//"path"
	//"bytes"
	"strings"
	"time"
)

func main() {

	log.Println("In main()")

	// Maintenance loop
	for {
		log.Println("In loop check and fix")

		// Maintenance functions
		postgresPatroniFix()

		time.Sleep(300 * time.Second)
	}
}

func postgresPatroniFix() {

	log.Println("In postgresPatroniFix()")

	pgStatefulSet := "mypostgres"

	//======================================
	// Get current state

	log.Println("Preparing command")

	cmd := fmt.Sprintf("kubectl exec -it -n default $(kubectl get pods -n default --no-headers -l %s-role=master | tail -n 1 | awk '{print $1}') -- bash -c \"cd /home/postgres && patronictl -c postgres.yml list\"", pgStatefulSet)

	log.Println("Running Command: " + cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		//log.Fatal(err)
		log.Println("Error running command: " + cmd)
		log.Println(err)
	}
	outputString := string(out)
	log.Println("Output is %s\n", outputString)

	//======================================
	// Get any errored member

	//cmd = fmt.Sprintf("kubectl exec -it -n default $(kubectl get pods -n default --no-headers -l %s-role=master | tail -n 1 | awk '{print $1}') -- bash -c \"cd /home/postgres && patronictl -c postgres.yml list\" | grep %s | awk -F'|' '!/running/ {print $3}' ", pgStatefulSet, pgStatefulSet)

	// tweaked command to just get 'running' items themselves instead of !running items
	cmd = fmt.Sprintf("kubectl exec -it -n default $(kubectl get pods -n default --no-headers -l %s-role=master | tail -n 1 | awk '{print $1}') -- bash -c \"cd /home/postgres && patronictl -c postgres.yml list\" | grep %s | awk -F'|' '/running/ {print $3}' ", pgStatefulSet, pgStatefulSet)

	log.Println("Running Command: " + cmd)
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		//log.Fatal(err)
		log.Println("Error running command: " + cmd)
		log.Println(err)
	}
	outputString = string(out)
	log.Println("Output string(out) is %s\n", outputString)

	// Reinit failed members
	failedMembers := strings.Split(outputString, "\n")
	for i, member := range failedMembers {

		// reinit if the string contains sts name (to avoid junk characters)
		if strings.Contains(member, pgStatefulSet) {
			log.Println("Doing patroni reinit for " + member + " - member id " + string(i))

			cmd = fmt.Sprintf("kubectl exec -n default -it $(kubectl get pods -n default --no-headers -l %s-role=master  | awk '{print $1}') -- bash -c \"cd /home/postgres && patronictl -c postgres.yml reinit %s %s --force \"", pgStatefulSet, pgStatefulSet, member)

			log.Println("Running Command: " + cmd)
			out, err = exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				//log.Fatal(err)
				log.Println("Error running command: " + cmd)
				log.Println(err)
			}
			outputString = string(out)
			log.Println("Output log.Println string(out) is %s\n", outputString)
		}
	}
}
