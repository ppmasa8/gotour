package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i % 15 == 0 {
			fmt.Println("fizzbuzz")
			break
		} else if i % 5 == 0 {
			fmt.Println("buzz")
			continue
		} else if i % 3 == 0 {
			fmt.Println("fizz")
			continue
		} else {
			fmt.Println(i)
			continue
		}
	}
}
