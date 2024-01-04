package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main()  {
	fmt.Println("Web requests in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close() // callers responsibilty to close the connection, otherwise it will be on 

	data, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(data)

	fmt.Println(content)


}