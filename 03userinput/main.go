package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the input")

	//comma ok synatax || comma error syntax
	// _ will have the error value
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}