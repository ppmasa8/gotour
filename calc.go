package main

import (
	"fmt"
	"strconv"
)

func calc(arg1 int, arg2 string, arg3 int) string {
	switch arg2 {
	case "+":
		return strconv.Itoa(arg1 + arg3)
	case "-":
		return strconv.Itoa(arg1 - arg3)
	case "*":
		return strconv.Itoa(arg1 * arg3)
	case "/":
		return strconv.Itoa(arg1 / arg3)
	default:
		return "Please type + - * / ."
	}
}

func main() {
	var arg1 int
	var arg2 string
	var arg3 int
	fmt.Println("Enter a formula　ex(29 * 9")
	fmt.Scan(&arg1, &arg2, &arg3)
	fmt.Println(calc(arg1, arg2, arg3))
}