//Using enum in

package main

import "fmt"

//Creating enum in golang
type Orderstatus int

//Iota in the enum is something that increment automatically
const (
	Received Orderstatus = iota
	Pending
	Success
	Failed
)

//Creating enum of type String
type studentStatus string

const (
	Good studentStatus = "good"
	Bad                = "bad"
	Poor               = "poor"
	Rich               = "rich"
)

func PrintOrderStatus(status Orderstatus) {
	fmt.Println("The Order status is : ", status)
}

func PrintStudentStatus(status studentStatus) {
	fmt.Println("The status of the student is : ", status)
}

func main() {
	PrintOrderStatus(Failed)
	PrintStudentStatus(Poor)
}
