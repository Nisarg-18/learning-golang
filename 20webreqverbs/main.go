package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Web requests in golang")

	// SendGetReq("http://localhost:8000/get")

	// PostRequest("http://localhost:8000/post")

	PostFormRequest("http://localhost:8000/postform")

}


func SendGetReq (url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	defer response.Body.Close()

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(content)) // one way

	// another way is to use strings package, done above
}

func PostRequest(url string) {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"go",
			"price":100,
			"platform":"YT"		
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PostFormRequest(uri string){
	//formdata
	data := url.Values{}
	data.Add("firstname", "nisarg")
	data.Add("lastname", "shah")
	data.Add("email", "nisarg@test.com")

	resp, err := http.PostForm(uri, data)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content,_ := io.ReadAll(resp.Body)

	fmt.Println(string(content))
}