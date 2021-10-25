package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in GO")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println(fruitList)

	var vegList = [3]string{"Tomato", "Potato", "Mushroom"}
	fmt.Println(vegList)
}