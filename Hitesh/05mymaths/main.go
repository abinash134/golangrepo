// Random no generation
package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	
)

func main(){
	//fmt.Println("welcome to math in golang")
	//var mynumOne int =1
	//var mynumTwo float64 = 4 

	//fmt.Println("The sum is:",mynumOne+int(mynumTwo))


	//Normal Random

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	//Random From Crypto
	
	myRandomNo, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(myRandomNo)

}