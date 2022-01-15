package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web request types")
	PerformPostformjsonRequest()
	//PerformGetRequest()
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

func PerformPostjsonRequest() {
	const url = "http://localhost:8000/post"

	requestbody := strings.NewReader(`
		{
			"name":"Abinash Pradhan"
		}
	`)

	response, _ := http.Post(url, "application/json", requestbody)

	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytelen, _ := responseString.Write(content)
	fmt.Println(bytelen)
	fmt.Println(responseString.String())

}
func PerformPostformjsonRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "abinsh")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytelen, _ := responseString.Write(content)
	fmt.Println(bytelen)
	fmt.Println(responseString.String())

}
