package main

import (
	"fmt"
	"log"
)

func main() {
	var myString string
	myString = "Green"
	log.Println("Logging string", myString)
	usePointer(&myString)
	fmt.Println("new value", myString)
}

func usePointer(s *string) {
	newVal := "Red"
	*s = newVal
}
