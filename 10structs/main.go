package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")

	// similar to classes in other languages
	// no inheritance in go, no parent or child

	nisarg := User{"Nisarg","nisarg@test.com",true,23}
	fmt.Println(nisarg)

	// placeholders in printf
	// %v for value, %+v for struct (prints key and value both)
	// %T for type

	fmt.Println(nisarg.Name)
}

type User struct{
	Name	string
	Email	string
	IsRegistered	bool
	Age	int
}