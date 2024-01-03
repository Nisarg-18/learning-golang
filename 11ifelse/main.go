package main

import "fmt"

func main()  {
	fmt.Println("If else in golang")

	loginCount := 9

	var result string

	if loginCount>10 {
		result = "More than 10"
	} else if loginCount<10 {
		result = "Less than 10"
	} else {
		result = "Equal to 10"
	}
	fmt.Println(result)
	
	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("False")
	}

	if num:=3; num<10 {
		fmt.Println("Num is less than 10")
	}

	// if err != nil{

	// }
}