package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GoLang")

	content := "This needs to be go in the file - Learning Files in Golang"
	file, err :=os.Create("./golearningfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./golearningfile.txt")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside file is: ", string(dataByte))
}