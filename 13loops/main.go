package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index, day:= range days{
		fmt.Printf("At %v day is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10{
		// if(rougueValue == 5){
		// 	break
		// }
		// if(rougueValue == 5){
		// 	rougueValue++
		// 	break
		// }
		// if(rougueValue == 5){
		// 	rougueValue++
		// 	continue
		// }

		
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
}