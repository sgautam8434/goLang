package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	personJson := `
	[{
"first_name" : "siddhant",
"last_name":"gautam"
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(personJson), &unmarshalled)
	if err != nil {
		log.Println("Error", err)
	}
	log.Println("unmarshalled: %v", unmarshalled)
	// write to json

	var secondPersom []Person

	var p1 Person
	p1.FirstName = "abhinav"
	p1.LastName = "Mallick"
	secondPersom = append(secondPersom, p1)

	secondJson, err := json.MarshalIndent(secondPersom, "", "  ")
	if err != nil {
		log.Println("error in creating json", err)
	}
	fmt.Println(string(secondJson))
}
