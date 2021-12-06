package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "Hello Web")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("no of bytes written :%d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "About Page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("no of bytes written :%d", n))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
