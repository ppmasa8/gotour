package main

import "fmt"

//(x, y int)ともできる
func add (x int, y int) int {
	return x+y
}

func main () {
	fmt.Println(add(34, 45))
}
