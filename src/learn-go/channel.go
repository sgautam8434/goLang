package main

import (
	"log"
	"math/rand"
)

const numPool = 10

func calculateValue(intChan chan int) {

	randomNumber := rand.Intn(numPool)
	intChan <- randomNumber
}
func main() {
	intChan := make(chan int)
	defer close(intChan)
	go calculateValue(intChan)
	num := <-intChan
	log.Println(num)
}
