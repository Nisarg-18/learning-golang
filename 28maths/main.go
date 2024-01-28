package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main(){
	fmt.Println("Maths in golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4

	// fmt.Println(numberOne+int(numberTwo)) // BAD

	// random number 

	// 1. math
	// rand.NewSource(time.Now().UnixNano())
	// fmt.Println(rand.Intn(6)+1) // 5 is not included

	// 2. crypto
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(6))

	fmt.Println(randomNumber)
}