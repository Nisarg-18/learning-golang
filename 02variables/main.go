package main

import "fmt"

// constants
// starts with capital letter -> public variable
const LoginToken string = "kasndskd"

func main()  {
	fmt.Println("Variables")
	var userName string  = "Nisarg";
	fmt.Println(userName)
	fmt.Printf("type of username is: %T \n", userName)

	var isLoggedIn bool  = true;
	fmt.Println(isLoggedIn)
	fmt.Printf("type of username is: %T \n", isLoggedIn)

	// uint8<uint16<uint32<uint64
	var intValue int  = 256;
	fmt.Println(intValue)
	fmt.Printf("type of username is: %T \n", intValue)

	// float32<float64
	var floatValue float32  = 256.234324343;
	fmt.Println(floatValue)
	fmt.Printf("type of username is: %T \n", floatValue)

	//default values
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("type of username is: %T \n", defaultInt)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("type of username is: %T \n", defaultString)

	//implicit type declaration
	var sample = "hello"
	fmt.Println(sample)
	// sample = 3 -> wrong

	// no var style (walrus operator, allowed only inside method)
	noOfUsers := 200
	fmt.Println(noOfUsers)
}