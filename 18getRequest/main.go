package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Https methods in golang")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "https://api.freeapi.app/api/v1/public/randomusers?page=1&limit=10"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	//Closing the request in Golang
	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	// fmt.Println(responseString.String()) --> this will print out all the reponse from the get url

}
