package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main(){
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}
