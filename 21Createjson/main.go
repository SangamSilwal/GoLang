package main

import (
	"encoding/json"
	"fmt"
)

// Creating the json Format
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
	//omitempty will not show the empty or nil value when we throw json
	//Using password `json:"-"` this will not take password in the json
}

func main() {
	fmt.Println("Json Handler in Golang")
	EncodeJson()

}

func EncodeJson() {
	newCourses := []course{
		{"ReactJs BootCamp", 299, "Youtube", "abc123", []string{"WebDev", "Python", "js"}},
		{"MongoDb BootCamp", 299, "Youtube", "abcasda123", []string{"WebDev", "Python", "js"}},
		{"Java BootCamp", 299, "Youtube", "sdaasda", []string{"WebDev", "Python", "js"}},
		{"SpingBoot BootCamp", 299, "Youtube", "asdad", []string{"WebDev", "Python", "js"}},
		{"Golang BootCamp", 299, "Youtube", "ssss", nil},
	}

	finalJson, err := json.Marshal(newCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", finalJson)

	//Using MarshelIndent
	modifiedJson, err := json.MarshalIndent(newCourses, "", "\t")
	//This will give me a good and modified json
	if err != nil {
		panic(err)

	}
	fmt.Printf("%s", modifiedJson)
	contentJson := string(modifiedJson)
	fmt.Printf("%T", contentJson)

}
