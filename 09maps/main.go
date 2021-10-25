package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["Go"] = "Go"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "PY")
	fmt.Println(languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}