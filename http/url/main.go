package main

import (
	"fmt"
	"net/url"
)

func main() {

	testing("http://localhost:8080")
	testing("http://uat.hdfcsky.com")
	testing("localhost:8080")

}

func testing(urlString string) {
	pu, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println(pu.String())
	fmt.Println("schema", pu.Scheme)
	fmt.Println("full", pu.Host)
	fmt.Println("hostname", pu.Hostname())
	fmt.Println("port", pu.Port())
	fmt.Println("---------------------")

}
