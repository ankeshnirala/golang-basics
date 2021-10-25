package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in GO")

	fruitList := []string{"Apple", "Peach"}

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 288
	highScores[2] = 256
	highScores[3] = 346
	// highScores[4] = 945

	highScores = append(highScores, 466, 233, 467)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)


	// how to remove a value from slices based on indexes
	courses := []string{"ReactJS", "NodeJS", "Python", "JavaScript"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}