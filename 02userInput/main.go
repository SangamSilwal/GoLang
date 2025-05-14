package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome To the Go ecosystem"
	fmt.Println(welcome)
	fmt.Println("Enter any Message: ")
	reader := bufio.NewReader(os.Stdin)

	// bufio  is the library which is used to read string from the user
	// which taking input from the stdin we should also take error too

	//comma ok syntx
	// input , err := reader.ReadString('\n)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("The Type of the input is %T\n", input)
}
