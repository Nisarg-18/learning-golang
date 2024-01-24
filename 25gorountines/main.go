package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // use a pointer
var mut sync.Mutex // use a pointer
var signals []string

func main () {
	// go greeter("Hello")
	// greeter("world")

	websiteList := []string {
		"https://go.dev",
		"https://lco.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}


	for _, web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // will always come at end, will tell the main to wait, coz all routines are not done yet
	fmt.Println(signals)
}

// func greeter (s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)		
// 	}
// }

func getStatusCode (endpoint string) {

	defer wg.Done() // should send done signal once it finishes
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
	}

	mut.Lock() //locks the resource, so only one go routine can access it, write in it, read it, while others cannot
	signals = append(signals, endpoint)
	mut.Unlock() // unlocks the resource
	fmt.Println(res.StatusCode)
}