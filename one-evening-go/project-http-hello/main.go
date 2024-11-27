package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, r *http.Request) {
	// Get the value of the `name` query parameter
	// localhost:8000/hello?name=Andrew
	name := r.URL.Query().Get("name")

	// Print a response message for this HTTP route
	fmt.Fprint(w, "Hello, ", name)
}

func main() {
	// Call the `helloFunction()` function when a request is sent to "/hello"
	http.HandleFunc("/hello", helloFunction)

	// http.ListenAndServe() starts an HTTP server on the `addr` address
	log.Fatal(http.ListenAndServe(":8080", nil))
}
