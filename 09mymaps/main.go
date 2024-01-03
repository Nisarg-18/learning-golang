package main

import "fmt"

func main()  {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	fmt.Println(languages["JS"])

	delete(languages, "JS")

	fmt.Println(languages)

	for key, value := range languages{
		fmt.Println(key,value)
	}

}