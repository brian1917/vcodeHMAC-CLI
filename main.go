package main

import (
	"fmt"
	"flag"
	"github.com/brian1917/vcodeHMAC"
)

var credsFile string
var targetUrl string
var method string

func init() {
	flag.StringVar(&credsFile, "credsFile", "", "Location of API credentials file")
	flag.StringVar(&targetUrl, "url", "", "URL")
	flag.StringVar(&method, "method", "", "HTTP Method")
}


func main() {
	flag.Parse()
	authHeader := vcodeHMAC.GenerateAuthHeader(credsFile, method, targetUrl)
	fmt.Println(authHeader)

}


