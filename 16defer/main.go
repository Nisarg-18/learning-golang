package main

import "fmt"

func main(){
	defer fmt.Println("Hello World") // defer means this will be executed at the end of the function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	// output -> Two then One then Hello World, LIFO order, bottom to top
	fmt.Println("Defer in golang")

	myDefer() 
	// Defer in golang, 4 3 2 1 0, Two, One, Hello World 
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 4 3 2 1 0
	}
}