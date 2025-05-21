/*
In Go encoding/Json package, NewEncoder and NewDecoder
provide stream oriented JSON encoding and decoding capabilities

##json.NewEncoder
Creates a new encoder that writes JSON to an output stream
(anything that implement io.writer)
io.Writer is an interface defined in the io package that represents the ability
to write data(usually bytes) to a destination(such as file,memory buffer, network connection)



##json.NewDecoder
Creates a new decoder that reads JSON from an input stream
input stream are anything that implement io.Reader

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// http.HandleFunc("/", handler)
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	fmt.Println("Error Starting server: ", err)
	// }
	file, _ := os.Open("output.json")
	defer file.Close()
	var config string
	json.NewDecoder(file).Decode(&config)
	fmt.Println(config)

}

// Writing Json to http response using NewEncoder
func handler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Hello world"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	//The above way is a Way to send json to the frontEnd
}

// Writing JSON to files using NewEncoder
func createFile() {
	file, _ := os.Create("output.json")
	defer file.Close()
	data := []string{"apple", "banana", "Cherry"}
	json.NewEncoder(file).Encode(data)
}
