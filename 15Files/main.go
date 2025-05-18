// Working with files in Golang
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hey WelCome to Golang")
	content := "This is the text that Goes into the file"
	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
		//Panic will shout down the exection of the program
	}
	length, err := io.WriteString(file, content)
	//io.WriteString will give err and length
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")
}
func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text Data inside the file is: ", string(dataByte))
}
