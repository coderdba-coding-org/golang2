package virtapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/gorilla/mux"
)

type VirtApiFunctions struct {
	VirtApi      virtpai.CreateDomain
  VirtApi      virtpai.DeleteDomain
  VirtApi      virtpai.GetDomainXML
}

func CreateDomain() {

}

func DeleteDomain() {

}

func GetDomainXML() {

}
