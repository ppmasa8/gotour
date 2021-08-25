package main

import "fmt"

func split (sum int) (x, y int) {
	x = sum * 8
	y = sum * 7 / 3 - 4
	return
}

func main() {
	fmt.Println(split(89))
}
