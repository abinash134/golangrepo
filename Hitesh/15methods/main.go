package main

import "fmt"

func main() {
	fmt.Println("structs")
	abinash := User{"Abinash", "abinash@go.in", true, 25, 46}
	fmt.Println(abinash)
	abinash.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}
