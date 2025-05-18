// Handling web request in Golang
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Checking Web Request")

	response, err := http.Get("https://hiteshchoudhary.com/index.html")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type: %T\n", response)
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(dataBytes))
}
