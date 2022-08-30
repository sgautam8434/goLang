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
	var1, var2 := returnTwoValues()
	fmt.Println("2 values", var1, var2)
}

func myFunc() string {
	return "myFunction"
}
func returnTwoValues() (string, string) {
	return "firstString", "secondString"
}
