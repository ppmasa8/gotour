package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main(){
	v := Vertex{1, 3}
	v.X = 5
	fmt.Println(v.X)
}
