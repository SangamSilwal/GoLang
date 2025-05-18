// Handling Urls in Golang
package main

import (
	"fmt"
	"net/url"
)

const Url string = "https://www.meresearch.org.uk/what-is-me?name=sangamsilwal&id=9083n"

func main() {
	fmt.Println("Welcome to handling Urls in Golang")
	result, _ := url.Parse(Url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//Using .Query is better way to parse a Url
	//Using Query gives a Map of the user

	qparam := result.Query()
	fmt.Println(qparam["id"][0])

	for _, val := range qparam {
		fmt.Println("Param is : ", val)
	}

	//We should always provide the  reference to the parts of url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "https://www.meresearch.org.uk",
		Path:    "/learn",
		RawPath: "user=Sangam",
	}

	//The parts of url return a complete URL
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
	fmt.Println(partsOfUrl)
}
