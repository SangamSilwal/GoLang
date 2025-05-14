//Type conversion in go
/*
the packages and their uses
-> bufio -> reading input from the stdin
-> strconv -> use fro string conversion
-> strings-> come out with lot of functionality
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Resturent Please Enter your Rating")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("The type of the input is: %T\n", input)
	convertedString, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The data Type of converted Data is %T\n", convertedString)
	}
}
