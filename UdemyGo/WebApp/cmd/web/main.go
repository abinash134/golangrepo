package main

import (
	"net/http"
	"webapp/pkg/handlers"
)

const portNumber = ":8080"
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
