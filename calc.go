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

func valid_iterator(arg2 string) string {
	var err = "Enter an +,-,*,/ in the second fields."
	iterator := [4]string{"+", "-", "*", "/"}
	if !arrayContains(iterator, arg2) {
		goto Warning
	}
	Warning:
		return err
}

func calc(arg1 int, arg2 string, arg3 int) (string, string) {
	if valid_int(arg1, arg3) != "" {
		goto Error
	}
	if valid_iterator(arg2) != "" {
		goto Error
	}
	switch arg2 {
	case "+":
		return strconv.Itoa(arg1 + arg3), ""
	case "-":
		return strconv.Itoa(arg1 - arg3), ""
	case "*":
		return strconv.Itoa(arg1 * arg3), ""
	case "/":
		return strconv.Itoa(arg1 / arg3), ""
	}
	Error:
		return valid_int(arg1, arg3), valid_iterator(arg2)
}

func main() {
	var arg1 int
	var arg2 string
	var arg3 int
	fmt.Println("Enter a formula　ex(29 * 9")
	fmt.Scan(&arg1, &arg2, &arg3)

	// For Debug
	fmt.Println(arg1, arg2, arg3)
	//

	fmt.Println(calc(arg1, arg2, arg3))
}

// Tools
func arrayContains(arr [4]string, str string) bool {
	for _, v := range arr{
		if v == str{
			return true
		}
	}
	return false
}