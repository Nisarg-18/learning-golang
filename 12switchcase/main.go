package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("switch case in golang")

	rand.NewSource(time.Now().UnixNano())

	randomNumber := rand.Intn(6) + 1

	fmt.Println(randomNumber)

	switch randomNumber {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough // will go to next case and run the code in the next case
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("default")
	}
}