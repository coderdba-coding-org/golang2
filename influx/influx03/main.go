package main

import (
        "github.com/gin-gonic/gin"
        cors "github.com/rs/cors/wrapper/gin"
	"log"
        "net/http"
	"encoding/json"
	"fmt"
        "os"
	"io/ioutil"
        "strings"
        "errors"
        client "github.com/influxdata/influxdb1-client/v2"

	//"time"
	//"errors"
        //"strconv"
        //"github.com/gin-contrib/cors"
        //client "github.com/influxdata/influxdb/client/v2"
)

