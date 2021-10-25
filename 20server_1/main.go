package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to first Server in GoLang")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	
	if(err != nil) {
		panic(err)
	}
	
	defer response.Body.Close()
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Learning GoLang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)

	response, _ := http.Post(myurl, "application/json", requestBody)

	defer response.Body.Close()
	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "ankesh")

	response, _ := http.PostForm(myurl, data)
	defer response.Body.Close()
	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}