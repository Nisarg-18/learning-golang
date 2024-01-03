package main

import "fmt"

func main()  {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Wednesday", "Saturday"}

	fmt.Println(days)

	// for loop
	// no ++i in go, only i++

	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index,day := range days{
	// 	fmt.Println(index,day)
	// }

	// similar to while loop
	rValue := 1
	for rValue < 10 {
		if rValue == 2{
			goto jump
		}
		if rValue == 5 {
			rValue++
			// break
			continue
		}
		fmt.Println(rValue)
		rValue++
	}

	// goto label
	jump:
		fmt.Println("Jump")
}