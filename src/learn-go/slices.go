package main

import "log"

func main() {
	var slice []string

	slice = append(slice, "val1")
	log.Println("slice", slice)
}
