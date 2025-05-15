/*
Syntax to declare array in GO
var array_name = [size]datatype{array Elements}

Decalaring array of undefined Size
array_name := [...]dataType{array Elemets}

Initializing array in GO
var array_name [array_size] datatype

### We can also initialize fix element in an array
arrayOfIntegers := [5]int{0: 7, 3: 9}
*/
package main

import "fmt"

func main() {
	student_name := [...]string{"Sangam", "Bhupal", "Dubey", "Pralad"}
	fmt.Println(student_name)

	student_marks := [4]int{1, 2, 3, 4}
	fmt.Println(student_marks)

	var myname [2]string
	myname[0] = "Sangam Silwal"
	myname[1] = "Subodh Silwal"
	fmt.Println(myname)

	array_of_number := [5]int{0: 1, 2: 3}
	fmt.Println(array_of_number)

}
