package main

import "fmt"

func main() {
	person1 := Person{"Sangam Silwal", 19, true, 98.97}
	fmt.Printf("The detail of the person is:\n %+v", person1)
	fmt.Printf("%T", person1)
	//Using +v make the result more detailed

	//initializing a struct instance
	var person2 Person
	person2 = Person{
		name:   "Dyana Roys",
		age:    45,
		isGood: false,
		marks:  54,
	}
	fmt.Printf("\n%+v", person2)

}

type Person struct {
	name   string
	age    int
	isGood bool
	marks  float32
}
type checkStatus func(int) bool

//We can also pass a function inside a structure
type Student struct {
	name   string
	age    int
	marks  float32
	status checkStatus
}
