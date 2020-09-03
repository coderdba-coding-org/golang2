package main

import (
    "fmt"
    "io/ioutil"
	"path/filepath"
	"encoding/json"

    "gopkg.in/yaml.v2"
)

type Service struct {
    APIVersion string `yaml:"apiVersion"`
    Kind       string `yaml:"kind"`
    Metadata   struct {
        Name      string `yaml:"name"`
        Namespace string `yaml:"namespace"`
        Labels    struct {
            RouterDeisIoRoutable string `yaml:"router.deis.io/routable"`
        } `yaml:"labels"`
        Annotations struct {
            RouterDeisIoDomains string `yaml:"router.deis.io/domains"`
        } `yaml:"annotations"`
    } `yaml:"metadata"`
    Spec struct {
        Type     string `yaml:"type"`
        Selector struct {
            App string `yaml:"app"`
        } `yaml:"selector"`
        Ports []struct {
            Name       string `yaml:"name"`
            Port       int    `yaml:"port"`
            TargetPort int    `yaml:"targetPort"`
            NodePort   int    `yaml:"nodePort,omitempty"`
        } `yaml:"ports"`
    } `yaml:"spec"`
}

func main() {

    filename, _ := filepath.Abs("./k1.yaml")
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
		//panic(err)
		fmt.Printf("Error reading file content")
    }

    //var config Config

    //err = yaml.Unmarshal(yamlFile, &config)
    //if err != nil {
    //    panic(err)
	//}
	//
	//fmt.Printf("Value: %#v\n", service.Firewall_network_rules)

	var service Service

    err = yaml.Unmarshal(yamlFile, &service)
    if err != nil {
		//panic(err)
		fmt.Printf("Error unmarshaling file content to struct")
	}
	
	fmt.Printf("Printing from Struct:\n")
    fmt.Printf("APIVersion value is: %#v\n", service.APIVersion)
	fmt.Printf("Kind value is: %#v\n", service.Kind)

	// Marshal to Json

	serviceJson, err := json.Marshal(&service)
	fmt.Printf("Printing from Json:\n")
	//fmt.Printf("Json is: %+v\n", serviceJson)
	fmt.Printf("%s\n", string(serviceJson))

}

