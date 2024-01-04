package main

import "fmt"

func main()  {
	fmt.Println("Functions in golang")
	greeter()
	result := adder(4,3)
	fmt.Println(result)

	result1, message := proAdder(1,2,3)

	fmt.Println(result1)
	fmt.Println(message)
}

func adder(valOne int, valTwo int) int {
	return valOne+valTwo
}

func proAdder(values ...int) (int,string){
	total := 0

	for _,val := range values{
		total += val
	}

	return total, "Hello"
}

// function inside a function is not allowed

func greeter(){
	fmt.Println("Hello")
}