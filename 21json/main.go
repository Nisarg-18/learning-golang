package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// aliases as 3rd parameter
	Name string `json:"coursename"` // will give key name as coursename in json response
	Price int // not compulsory to define the third parameter
	Platform string `json:"website"`
	Password string `json:"-"` // will not come up in json, will be ignored
	Tags []string `json:"tags,omitempty"` // omitempty will skip it if it is nil
}

func main()  {
	fmt.Println("Welcome to json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson () {
	courses := []course{
		{"GO", 299, "YT", "abc", []string{"go"}},
		{"MERN", 199, "YT", "123", []string{"web-dev", "js"}},
		{"Python", 499, "YT", "abc123", nil},
	}

	//convert this data into JSON
	jsonData, err := json.MarshalIndent(courses,"","\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
			"coursename": "GO",
			"Price": 299,
			"website": "YT",
			"tags": ["go"]
		}
	`)

	var myCourse course

	checkJson := json.Valid(jsonDataFromWeb)

	if checkJson {
		json.Unmarshal(jsonDataFromWeb, &myCourse)

		fmt.Printf("%#v\n",myCourse)
	} else {
		fmt.Println("JSON not valid!")
	}	

	// if we don't want data to be stored in form of struct, we can use maps

	var data map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &data)

	fmt.Printf("%#v\n", data)

	for k,v := range data {
		fmt.Println(k,v)
	}
}