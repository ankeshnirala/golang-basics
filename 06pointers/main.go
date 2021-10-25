package main

import "fmt"

func main() {
	fmt.Println("Welcome to go pointers")

	// var myNumber *int
	// fmt.Println(myNumber)

	myNumber := 23
	ptr := &myNumber

	fmt.Println(ptr, *ptr)

	*ptr = *ptr + 4
	fmt.Println(myNumber)
}