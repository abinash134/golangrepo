package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hello file")
	content := "This id the first file intraction in golang"

	file, err := os.Create("./abinsh.txt")

	if err != nil {
		panic(err)
	}
	lenght, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("the lenght is", lenght)
	defer file.Close()
	readFile("./abinsh.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data is\n", string(databyte))
}
