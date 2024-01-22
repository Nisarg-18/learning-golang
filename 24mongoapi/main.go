package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main()  {
	fmt.Println("Welcome to mongo api")
	r := router.Router()

	fmt.Println("Starting server")

	log.Fatal(http.ListenAndServe(":8000",r))

	fmt.Println("Server running")
}