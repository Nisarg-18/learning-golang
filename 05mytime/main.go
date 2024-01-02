package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2023,time.October,18,23,23,23,23,time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}