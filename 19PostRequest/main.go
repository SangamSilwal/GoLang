// Post Requests in Golang
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "SampleUrl"
	requestBody := strings.NewReader(`
	{
	"CourseName":"Learn Goalang with sangam Silwal",
	"Price":0,
	"AvailableStatus":true
	}
	`)
	fmt.Println(requestBody)
	fmt.Printf("\nThe dataTye of using .NewReader is : \n%T", requestBody) //The output: *strings.Reader

	response, err := http.Post(myurl, "application/json", requestBody)
	//The above is the syntax for performing Post in Golang
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // Closing the server using the defer

	content, _ := io.ReadAll(requestBody)
	fmt.Println(string(content))

}
