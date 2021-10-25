package main

import (
	"fmt"
	"net/url"
)

const myurl string= "https://lco.dev:3000/learn?coursename=golang&paymentid=asdnsfksn"

func main() {
	fmt.Println("URLs in GoLang")

	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Println(qparams)
	for key, value := range qparams{
		fmt.Println(key, value)
	}

	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=ankesh&tutid=123",
	}

	anotherUrl := partsOfUrl.String();
	fmt.Println(anotherUrl)
}