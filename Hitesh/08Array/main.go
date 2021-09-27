package main

import "fmt"

func main(){
	fmt.Println("Array")
	//declaring a array

	var fruitList [3]string

	fruitList[0]  = "Orange"
	fruitList[1] = "Banana"
	fruitList[2] = "Peach"


	var Veg = [3]string{"potato","tomato","cucumber"}

	fmt.Println(Veg)

	fmt.Println(fruitList)
}