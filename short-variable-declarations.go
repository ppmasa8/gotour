package main

import "fmt"

func main() {
	var i, j int = 2, 4
	//関数の中でのみ:=がvarと同様な仕様で使える。
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, k, python, java)
}
