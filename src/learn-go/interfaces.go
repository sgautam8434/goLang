package main

import "fmt"

type Animal interface {
	Sys() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Color string
}

func main() {

	dog := Dog{"labra", "labrador"}
	PrintInfo(&dog)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Sys(), "and has", a.NumberOfLegs(), "legs")
}
func (d *Dog) Sys() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
