package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

// Greet says hello for someone
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler handles requests from API
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
