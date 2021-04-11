package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/status", helloWorld)
	http.ListenAndServe(":5000", nil)
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	_, _ = fmt.Fprintf(writer, "<h1>Hello! It's up and running</h1>")
}
