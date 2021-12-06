package main

import "fmt"

func main() {
	fmt.Println("structs")
	abinash := User{"Abinash", "abinash@go.in", true, 25}
	fmt.Println(abinash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
