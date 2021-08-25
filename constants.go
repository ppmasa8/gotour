package main

import "fmt"

const Pi = 3.14

func main() {
	//定数は:=では表せない
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

