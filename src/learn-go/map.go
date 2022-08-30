package main

import (
	"log"
)

type Users struct {
	firstName string
	lastName  string
}

func main() {
	myMap := make(map[string]string)
	myMap["name"] = "Siddhant"
	log.Println(myMap["name"])
	mapWithStruct := make(map[string]Users)
	mapWithStruct["name"] = Users{"Siddhant", "Gautam"}
	log.Println("full name", mapWithStruct["name"])
}
