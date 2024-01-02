package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple","Banana"}

	fmt.Printf("type %T\n", fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println(fruitList)

	//slice the slices
	// start is inclusive and end is non inclusive
	// start:end -> format
	fruitList = append(fruitList[1:2])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0]=123
	highScores[1]=234
	highScores[2]=456
	highScores[3]=890
	// highScores[4]=1234

	highScores = append(highScores, 100)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	// remove a value from slice based on its index
	var courses = []string{"php", "go", "react"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}