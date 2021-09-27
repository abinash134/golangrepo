package main

import "fmt"

func main(){
	fmt.Println("Pointers")
	//Declaring a pointer
	var myPtr *int
	fmt.Println("my pointer is ",myPtr)
	//asigning values to the pointers
	var myVar  int = 23
	var myNewPtr = &myVar
	fmt.Println(*myNewPtr)

}