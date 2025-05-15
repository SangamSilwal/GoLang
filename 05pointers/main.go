package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on pointers")

	//initializing pointer in go
	// var ptr *int
	// fmt.Println((ptr)) --> It will return nil as it not pointing to anything
	mynumber := 23
	var ptr = &mynumber

	//This will print the memory location of the mynumber
	// fmt.Println(ptr)
	// &myNumber is the reference of mynumber

	fmt.Println("Memory Address: ", ptr)
	fmt.Println("Value pointer holding is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println(mynumber)
}
