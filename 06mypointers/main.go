package main

import "fmt"

func main()  {
	// pointers gaurantees the value of a variable will remain same everywhere
	// as address of the variable is used
	fmt.Println("Welcome to pointers")

	// var ptr *int // pointer decalaration
	// fmt.Println(ptr) // nil

	// * -> pointer 
	// & -> reference
	myNumber := 23

	ptr := &myNumber // ptr will have address of myNumber

	fmt.Println(ptr) // address 
	fmt.Println(*ptr) // actual value 23

	*ptr = *ptr + 2 // ptr has address, *ptr gives value stored at that address
	fmt.Println(*ptr) // 25

}

