package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("Files in golang")

	content := "file content"

	file, err := os.Create("./sample.txt")

	checkNilError(err)

	length, err := io.WriteString(file,content)
	checkNilError(err)

	fmt.Println(length)

	defer file.Close()

	readFile("./sample.txt")

}

func readFile (fileName string) {
	data, err := os.ReadFile(fileName)

	checkNilError(err)

	fmt.Println(string(data))
}

func checkNilError (err error) {
	if err != nil {
		panic(err)
	}
}