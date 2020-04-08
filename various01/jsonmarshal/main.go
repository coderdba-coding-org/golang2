package main

import (
	"encoding/json"
	"fmt"
	//"encoding/xml"
	//"strings"
)

type VmImageRequest struct {
	ImageName string `json:"imageName"`
	MD5Sum    string `json:"md5sum"`
}

func main() {

	// create the request body as structure
	vmImageRequest := VmImageRequest{ImageName: "centos7base", MD5Sum: "ljkdsf09823kjnsf8923"}

	// create a json payload from the structure
	payload, err := json.Marshal(&vmImageRequest)
	if err != nil {
		fmt.Printf("Vm Image Request: JSON marshalling failed")
		//return 0, fmt.Errorf("Vm Image Request: JSON marshalling failed")
	}

        fmt.Printf("%+v\n", vmImageRequest)
        fmt.Printf("%s\n", string(payload))

}
