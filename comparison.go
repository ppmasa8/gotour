package main

import (
	"fmt"
	"math/rand"
	"os"
)

func genNum() int {
	// Generate rand
	n := rand.Intn(100)
	return n
}

func main() {

	// Player number
	var a int
	fmt.Println("Please type a number 0~100.")
	fmt.Scan(&a)

	// Validation
	if a < 0 || a > 100 {
		fmt.Println("You can input 0~100.")
		os.Exit(1)
	}

	// CPU number
	b := genNum()

	// Log
	fmt.Println("Your number is", a)
	fmt.Println("Cpu number is", b)


	// Result
	if b < a {
		fmt.Println("Win")
	} else if b > a {
		fmt.Println("Lose")
	} else {
		fmt.Println("Draw")
	}
}
