package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name		string 		`json:"coursename`
	Price		int		
	Platform 	string		`json:"website"`
	Password 	string		`json:"-"`
	Tags 		[]string	`json:"tags,omitempty`
}

func main() {
	fmt.Println("Welcome to a Bit more JSON in GoLang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Node Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "lco.in", "abc123", nil},
	}

	// package this data as JSON data
	// finalJson, _ := json.Marshal(lcoCourses)
	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "Name": "React Bootcamp",
                "Price": 299,
                "website": "lco.in",
                "Tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was vlaid!!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}