package main

import "fmt"

func main() {
	fmt.Println("Structs in GOLANG")

	// no inheritance in golang
	// no super or parent in golang

	ankesh := User{"Ankesh", "ankesh.go.dev", true, 24}
	fmt.Println(ankesh)
	fmt.Printf("Ankesh details are: %+v\n", ankesh)
	fmt.Printf("Name is %v and Email is %v\n", ankesh.Name, ankesh.Email)
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}