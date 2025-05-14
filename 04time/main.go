package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	//01-02-2006 this is the format
	//---> format 01-02-2006 15:04:05 Monday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	//The output will be the present time

	createdDate := time.Date(2024, time.August, 10, 23, 23, 0, 0, time.UTC)
	//year month date hour minute sec
	fmt.Println(createdDate)
}
