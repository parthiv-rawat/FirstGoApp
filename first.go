package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Print("The server is running at port 8080.")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(write http.ResponseWriter, read *http.Request) {
	fmt.Fprint(write, "Hello, World!")
}
