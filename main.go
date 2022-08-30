package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var stringData string
	stringData = "Stringdata"
	fmt.Println(stringData)
	var i = 8
	fmt.Println("interger", i)
	myFunctionReturnVal := myFunc()
	fmt.Println("call function", myFunctionReturnVal)
}

func myFunc() string {
	return "myFunction"
}
