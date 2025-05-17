/*
Functions in GO

--> Anonymous Function in Go
Anonymous Function in Go are the function
without the function name

For using Anonymous Function:= We assign the function
to a variable and use the function

Example of Anonymous Function:=
var name = func()
{
fmt.Println("Hey this is Anonymous Function")
}

*/

package main

import "fmt"

func main() {
	SampleFunction()
	result, message := ProAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(message, result)

	//Making Anonymous Function
	var adder = func(n1 int, n2 int) {
		fmt.Println(n1 + n2)
	}
	adder(2, 3)

}

func SampleFunction() {
	fmt.Println("Hey this is from Sample Function")
	sum, product := AddAndMultiply(2, 3)
	fmt.Printf("Sum is : %v and product is %v", sum, product)

}

//Function return type with parameter in go
func AddAndMultiply(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	product := n1 * n2
	return sum, product

}

//Pro adder function in goLang
func ProAdder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "The total sum is returned"
}

//Function Which returns an Anonymous Function
func Returner(n1 int) func() int {
	return func() int {
		n1++
		return n1
	}

}
