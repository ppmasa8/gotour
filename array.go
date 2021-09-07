package main

import "fmt"

func main(){
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 4, 6, 3, 6, 8}
	fmt.Println(primes)
}
