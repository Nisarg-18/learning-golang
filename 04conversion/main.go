package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome")
	fmt.Println("Please input your rating between 1 and 5")
	
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("You rated", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("added", numRating+1)
	}
}