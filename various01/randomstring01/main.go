package main

// https://www.calhoun.io/creating-random-strings-in-go/

import (
  "math/rand"
  "time"
  "fmt"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}


func CreateRandomString(length int) string {

const charSet = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seedRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

  b := make([]byte, length)
  for i := range b {
    b[i] = charSet[seedRand.Intn(len(charSet))]
  }
  return string(b)
}

func main() {

fmt.Println(StringWithCharset(5, charset))
fmt.Println(String(7))
fmt.Println(CreateRandomString(9))

name := "origname"
fmt.Println(name + "-" + CreateRandomString(5))

}
