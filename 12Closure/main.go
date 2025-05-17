/*
Closure is a nested functionthat allow us to access
variables of the outer function even after the outer function
is closed



while importing packages from other file we can import like
import _ "packages/pac_name" ---> This will help us
to get out of error even if the package is not used
*/

package main

import "fmt"

func main() {
	name := greet()
	fmt.Println(name())

	num1 := printOddNumber()
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())
}

func greet() func() string {
	name := "Sangam Silwal"
	return func() string {
		return name
	}
}

func printOddNumber() func() int {
	number := 1
	return func() int {
		number = number + 1
		return number
	}
}
