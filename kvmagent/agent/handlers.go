package agent

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

type Handlers struct {
	VirtApi      virtapi.VirtApiFunctions
	ImageHandler imagehandler.ImageHandlers
}
