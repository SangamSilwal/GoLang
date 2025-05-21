package main

import "fmt"

//Using generic type in Function
//We can also use [T interference{}]
func printSlice[T any](items []T) {
	for _, value := range items {
		fmt.Println(value)
	}
}

//Limiting the type used in the function
func printLimitedSlice[T int | string](item []T) {
	for _, value := range item {
		fmt.Println(value)
	}
}

//Using Generic with struct
type stack[T any] struct {
	element []T
}

type quene[T comparable, V string] struct {
	element        []T
	anotherElement V
}

func main() {
	name := []string{"sangam", "Silwal"}
	printSlice([]int{1, 2, 3})
	printSlice(name)
	printLimitedSlice([]int{1, 2, 3, 4, 5})

	newStack := stack[string]{
		element: []string{"sangam", "silwal", "is", "my", "name"},
	}
	fmt.Println(newStack)

}
