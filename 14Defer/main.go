package main

import "fmt"

func main() {
	//Defer works as last in First out
	defer fmt.Println("World")  //This will run at the last of the function
	defer fmt.Println("World1") //This will run at the last of the function
	defer fmt.Println("World2") //This will run at the last of the function
	fmt.Println("Hello")

	myDefer()
}

//the above output will be world2, world1 and world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
