package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	name := []string{"gautam", "siddhant", "abhinav", "mallick"}
	for _, name1 := range name {
		log.Println(name1)
	}
}
