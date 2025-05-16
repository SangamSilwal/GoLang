/*
Map are the dataStructures that store
key/Value
Example of using map:
syntax: map[string]int{key: value}
*/
package main

import "fmt"

func main() {
	subMarks := map[string]int{"Sangam": 1, "Silwal": 2, "Subham": 3, "Dogle": 4}
	fmt.Println(subMarks)
	fmt.Printf("The data Type of Map Object is : %T\n", subMarks)

	//Initializing map using make
	studentSheet := make(map[string]int)
	studentSheet["Sangam"] = 97
	studentSheet["Sugam"] = 99
	fmt.Println(studentSheet)

	//updating map
	studentSheet["Sugam"] = 87
	fmt.Println(studentSheet)

	//Deleting map
	delete(studentSheet, "Sangam")
	fmt.Println("Deleted Map: ", studentSheet)

	//Looping thorugh a map in GoLang
	for key, value := range subMarks {
		fmt.Printf("The key is %v and the Value is %v\n", key, value)
	}

	//using comma ok syntax for accesing key only
	for key, _ := range subMarks {
		fmt.Printf("The key is %v\n", key)
	}

}
