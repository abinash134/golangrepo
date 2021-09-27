package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("its time")
	presenttime := time.Now()
	fmt.Println(presenttime)
	//formatting syntax is this and constant and cant be channged
	fmt.Println(presenttime.Format("01-03-2006 15:04:06 monday"))

	//creating a date manually
	createdate := time.Date(2021 , time.June,21,20,48,56,0,time.Local)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-02-2006 Monday"))
}
