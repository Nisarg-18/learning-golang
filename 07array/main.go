package main

import "fmt"

func main()  {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Pineapple"
	// fruitList[0] = "Apple"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato","mushroom"}
	fmt.Println(len(vegList))
}