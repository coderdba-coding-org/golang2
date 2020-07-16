package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type Tracks struct {
    Toptracks Toptracks_info
}

type Toptracks_info struct {
    Track []Track_info
    Attr  Attr_info `json: "@attr"`
}

type Track_info struct {
    Name       string
    Duration   string
    Listeners  string
    Mbid       string
    Url        string
    Streamable Streamable_info
    Artist     Artist_info   
    Attr       Track_attr_info `json: "@attr"`
}

type Attr_info struct {
    Country    string
    Page       string
    PerPage    string
    TotalPages string
    Total      string
}

type Streamable_info struct {
    Text      string `json: "#text"`
    Fulltrack string
}

type Artist_info struct {
    Name string
    Mbid string
    Url  string
}

type Track_attr_info struct {
    Rank string
}
