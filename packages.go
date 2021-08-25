package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Println(math.pi)だとエクスポートされていないため、エラーになる
	fmt.Println(math.Pi)
}

