package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func valid_int(arg1 int, arg3 int) string {
	var err = "Enter an integer in the first and third fields."
	if reflect.TypeOf(arg1).Kind() != reflect.Int || reflect.TypeOf(arg3).Kind() != reflect.Int {
		goto Warning
	}
	Warning:
		return err
}

func calc(arg1 int, arg2 string, arg3 int) string {
	if valid_int(arg1, arg3) != "" {
		goto Error
	}
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
	Error:
		return valid_int(arg1, arg3)
}

func main() {
	var arg1 int
	var arg2 string
	var arg3 int
	fmt.Println("Enter a formula　ex(29 * 9")
	fmt.Scan(&arg1, &arg2, &arg3)
	fmt.Println(arg1, arg2, arg3)
	fmt.Println(calc(arg1, arg2, arg3))
}