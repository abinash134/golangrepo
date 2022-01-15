package main

import "fmt"
func main() {
	ch := make(chan int)
	ch <- 7
	val := <-ch
	fmt.println(val)
}