// Random number guessing game in go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// For testing until I learn how to accept input
const maxrand = 20

func setrandno(max int) (a int) {
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(max) + 1
	return a
}

func main() {
	fmt.Println("The Random Number Guessing Game in Go")
	randno := setrandno(maxrand)
	fmt.Println("Random number is", randno)
}
