package imagehandler


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

type ImageHandlers struct {
	VirtApi      virtpai.DownloadImage
  VirtApi      virtpai.CheckImageExistsLocally
}

func DownloadImage() {

}

func CheckImageExistsLocally() {

}
