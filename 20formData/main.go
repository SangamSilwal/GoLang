package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	sendFormData()
}

func sendFormData() {
	data := url.Values{}
	data.Add("FirstName", "Sangam")
	data.Add("LastName", "Silwal")
	data.Add("email", "sangamsilwal@gmail.com")

	fmt.Println(data)
	fmt.Printf("The type of Data is : %T \n", data)

	//Using Post Form to post the data
	response, err := http.PostForm("randomUrl", data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
