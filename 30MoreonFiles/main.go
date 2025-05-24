package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	//reading File
	// f, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// buf := make([]byte, 12)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("data", string(buf), d)

	// f, err := os.ReadFile("hello.txt")
	//Read file is good for smaller files
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	//Read Folders
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// fileinfo, err := dir.ReadDir(-1)

	// for _, fi := range fileinfo {
	// 	fmt.Println(fi.Name())
	// }

	sourceFile, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	//Making destFile
	destfile, err := os.Create("hello2.txt")
	if err != nil {
		panic(err)
	}
	defer destfile.Close()

	//Making newReader and newWriter
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destfile)

	for {
		byteRead, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		fmt.Println("The bytes Read are: ", string(byteRead))
		er := writer.WriteByte(byteRead)
		if er != nil {
			panic(er)
		}

	}
	writer.Flush()
}
