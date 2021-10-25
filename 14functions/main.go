package main

import "fmt"

func main() {
	greeter()
	result := adder(4, 6)
	fmt.Println("Welcome to functions in GoLang")
	fmt.Println("Result is: ", result)

	proResult, proMsg := proAdder(4, 5, 6)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println(proMsg)
}

func proAdder(values ...int) (int,string) {
	total := 0

	for _, value := range values{
		total += value
	}

	return total, "Hello I am from pro adder"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namstey from GoLang")
}

