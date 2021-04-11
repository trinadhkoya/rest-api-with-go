package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":5001", nil)
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello! Go...")
}
