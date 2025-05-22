package main

import (
	"fmt"
	"sync"
)

/*
Channels in goroutine act as a medium to communicate with each other

Basically goroutines are use to make concurrent programs. Concurrent programs are those
programs that runs parallely. While running parallely they need a meduim to communicate with
each other. Go channels is way for the goRoutine to communicate with each Other
*/

func sendDataToChannel(number chan int, stringValue chan string, w *sync.WaitGroup) {
	defer w.Done()
	number <- 15
	stringValue <- "Learning Go channel"
	fmt.Println("Data sent to the Goroutine using Golang Channels")
}

func makeSum(result chan int, num1 int, num2 int) {
	newSum := num1 + num2
	result <- newSum
}

func main() {

	//Creating a channel in golang
	var wg sync.WaitGroup
	number := make(chan int)
	stringValue := make(chan string)
	wg.Add(1)
	go sendDataToChannel(number, stringValue, &wg)
	fmt.Println("The Number value is : ", <-number)
	fmt.Println("The string value is : ", <-stringValue)

	newNumber := make(chan int)
	go makeSum(newNumber, 1, 2)
	finalResult := <-newNumber
	fmt.Println("the sum of the numbers is: ", finalResult)

}
