package main

import "fmt"

func main() {
	fmt.Println("hello world")
	defer fmt.Println("hello defer")
	fmt.Println("by defer")
	defer Deferfunc()
	fmt.Println("by deferssss")

}

func Deferfunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
