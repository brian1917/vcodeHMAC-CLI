package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/brian1917/vcodeHMAC"
)

var credsFile string
var targetURL string
var method string

func init() {
	flag.StringVar(&credsFile, "credsFile", "", "Location of API credentials file")
	flag.StringVar(&targetURL, "url", "", "URL")
	flag.StringVar(&method, "method", "", "HTTP Method")
}

func main() {
	flag.Parse()
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, method, targetURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(authHeader)

}
