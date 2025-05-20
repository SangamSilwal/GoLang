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
	DecodeJson()

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

func DecodeJson() {
	//Json data from Web is in bytes
	jsonDataFromWeb := []byte(`
	            {
                "coursename": "SpingBoot BootCamp",
                "Price": 299,
                "Platform": "Youtube",
                "Password": "asdad",
                "tags": ["WebDev","Python","js"]
				}
	`)
	fmt.Println(jsonDataFromWeb) //This will print the jsonDataFromWeb in byte format
	var courseStruct course
	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &courseStruct)
		fmt.Printf("%#v\n", courseStruct)
	} else {
		fmt.Println("The json was not valid")
	}

	//some cases where you just want to add data to key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println(myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key : %v and value: %v\n", key, value)
	}

}
