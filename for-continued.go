package main

import "fmt"

func main() {
	sum := 1
	//インクリメントと定数てきなものは省略可能
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
