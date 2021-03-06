// Random number guessing game in go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// For testing until I learn how to accept input
const maxrand = 2

func setrandno(max int) (a int) {
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(max) + 1
	return a
}

func checkguess(answer, guess int) (counter int) {
	counter = 1
	for guess != answer {
		counter += 1
		switch {
		case guess <= 0:
			fmt.Println("Your answer is far too small. The number lies between 1 and", maxrand)
		case guess > maxrand:
			fmt.Println("Your answer is far too big. The number lies between 1 and", maxrand)
		case guess < answer:
			fmt.Println("Too small. Try something bigger.")
		case guess > answer:
			fmt.Println("Too big. Try something smaller.")
		}
	}
	return counter
}

func main() {
	fmt.Println("The Random Number Guessing Game in Go")
	randno := setrandno(maxrand)
	fmt.Println("Random number is", randno)
	fmt.Println("You guessed in", checkguess(randno, 1), "go(es)")
}
