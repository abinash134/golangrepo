package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web request types")
	PerformGetRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"
	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()
	fmt.Println(responce.StatusCode)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	bytelen, _ := responseString.Write(content)
	fmt.Println(bytelen)
	fmt.Println(responseString.String())
	//fmt.Println(string(content))
}
