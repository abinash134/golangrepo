package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://lco.dev")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Of type : %T \n", resp)
	defer resp.Body.Close()

	databyte, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(databyte))
}
