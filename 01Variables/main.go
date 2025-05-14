package main

import "fmt"

// jwtToken := 50000 ---> This will give error
//Outside method := is not allowed

const LoggedInToken = "SangamSiwal123"

//Using L -> Capital will make the consant public and can be access anywhere from the focument

func main() {
	var username string = "Sangam Silwal"
	fmt.Println(username)

	//This is used for finding the type of the variable
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Printf("Varibel is of type : %T \n", isLoggedIn)

	var anotherInt int
	fmt.Println(anotherInt)
	fmt.Printf("Varibale is of type: %T \n", anotherInt)

	//implicit type
	var website = "Learnwith ee"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)
}
