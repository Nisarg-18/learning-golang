package main

import "fmt"

// methods are nothing but functions inside structs (classes)

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

	nisarg.GetStatus()

	nisarg.NewMail() // new@test.com // change the email in copy of struct, not the original one

	fmt.Println(nisarg.Email) // nisarg@test.com // has the original value, as a copy is passed in NewMail()
}

type User struct{
	Name	string
	Email	string
	IsRegistered	bool
	Age	int
	// oneAge int //private 
	// small letter starting -> not exported, private
	// capital letter starting -> exported, public
}

// method 

func (u User) GetStatus(){
	fmt.Println(u.IsRegistered)
}

func (u User) NewMail(){
	u.Email = "new@test.com"
	fmt.Println(u.Email)
}