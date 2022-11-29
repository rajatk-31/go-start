package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello from file ssave")
	content := "This will go in the file"

	file, err := os.Create("./test.txt")
	checkError(nil)

	length, err := io.WriteString(file, content)
	defer file.Close()
	checkError(err)
	fmt.Println("File length is : ", length)
	readFile("./test.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkError(err)

	fmt.Println("Text data in file is : ", string(dataByte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
