package main

import "fmt"

func main() {
	defer myDefer()
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func myDefer() {
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}