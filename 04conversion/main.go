package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Pizza App"
	fmt.Println(welcome)
	fmt.Println("Please rate our pizza app between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to your rating, ", numRating+1)
	}

}