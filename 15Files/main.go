// Working with files in Golang
package main

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	fmt.Println("Hey WelCome to Golang")
// 	content := "This is the text that Goes into the file"
// 	file, err := os.Create("./myFile.txt")
// 	if err != nil {
// 		panic(err)
// 		//Panic will shout down the exection of the program
// 	}
// 	length, err := io.WriteString(file, content)
// 	//io.WriteString will give err and length
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Length is: ", length)
// 	defer file.Close()
// 	readFile("./myFile.txt")
// }
// func readFile(filename string) {
// 	dataByte, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Text Data inside the file is: ", string(dataByte))
// }

func main() {
	fmt.Println("File Handling in Golang")
	context := "This is to be written in Golang using file Handling"
	Createfile("./hello.txt", context)

}
func Createfile(filename string, context string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(file, context)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
