package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=abcbd"

func main()  {
	fmt.Println("Handling URLs in golang")

	fmt.Println(myUrl)

	// URL to parts
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()

	for key, value := range qParams{
		fmt.Println(key,value)
	}

	// parts to URL
	partsToUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=nisarg",
	}

	newUrl := partsToUrl.String()

	fmt.Println(newUrl)
}