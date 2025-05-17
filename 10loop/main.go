// loop in go
package main

import "fmt"

func main() {
	daya := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thurday"}
	for index, value := range daya {
		fmt.Println("The index is : ", index)
		fmt.Println("The value is: ", value)
	}

	for _, value := range daya {
		fmt.Printf("%v", value)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 3 {
			goto here
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
here:
	println("Hey this is here")

}
