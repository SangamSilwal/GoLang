package main

import (
	"fmt"
	"sync"
)

/*
In golang we use goroutine to create concurrent environment
concurrent program are able to run multiple program at the same time
goroutine are used in golang to sun two or more than two function asychronously
*/

//using the concept of waitGroup in golang

func runGo(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(number)
}

func main() {
	var waitGp sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGp.Add(1)
		fmt.Println("This is Goalng goroutine", i)
		go runGo(i, &waitGp)
	}

}

//To use Goroutine in golang we need to use go waitgroup
//Just remember Waitgroup is a package found in sync function
//To use waitGroup we need to pass reference to the function
//We also need to use Add and .Done() method in waitGroup
