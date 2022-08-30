package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthDate   time.Time
}

func main() {

	user := User{firstName: "Siddhant", lastName: "Gautam", phoneNumber: "9099879093"}
	fmt.Println(user.lastName)
}
