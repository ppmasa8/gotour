package main

import "fmt"

func swap (x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("中井貴一", "中居くん")
	fmt.Println(a, b)
}