/*
Slice are like array in Go
But the value of slice is not defined
It is much used in Go than array

-> Declaring a Slice in go
numbers := []int{numbers element}


->functions in Slice
Functions	Descriptions
append()	adds element to a slice
copy()	copy elements of one slice to another
Equal()	compares two slices
len()	find the length of a slice
*/

package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	fmt.Println(numbers[1:3])

	//Apending in a Slice
	numbers = append(numbers, 11, 12, 13)
	fmt.Println(numbers)

	//Creating Slice using make
	highScore := make([]int, 3)
	highScore[0] = 1
	highScore[1] = 2
	highScore[2] = 3
	// highScore[3] = 8 --> This will give an error because we have decalare the size to 3

	highScore = append(highScore, 45, 67, 4, 6, 9, 90) //We can use this to append value inside the slice
	//It allocate memory again fot new elements

	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println("Sorted Slice are: ", highScore)

	//To check if it is sorted or not we can use
	fmt.Println(sort.IntsAreSorted(highScore))

	//Making a slice of string
	newNames := make([]string, 3, 4)
	newNames[0] = "Sangam"
	newNames[1] = "Silwal"
	newNames[2] = "Subham"
	newNames = append(newNames, "Joseph", "Pandit")
	sort.Strings(newNames)
	fmt.Println(newNames)

	//----> Different Ways to Delete an element from a slice in Go
	//deleting silwal
	newNames = append(newNames[:3], newNames[4:]...)
	fmt.Println("Using append: ", newNames)

	//using Slices.Delete
	newNames = slices.Delete(newNames, 1, 4)
	fmt.Println(newNames)

	//Using Slices.DeleteFunc
	highScore = slices.DeleteFunc(highScore, func(s int) bool {
		return s%2 != 0
	})
	fmt.Println("HighScore Using DeleteFunc only even", highScore)

	fmt.Println(slices.Index(newNames, "silwal"))

}
