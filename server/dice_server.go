package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", DiceHandler) // Register a handler
	http.ListenAndServe(":8888", nil) // Start the server
}

// DiceHandler  Describes what the server does.
func DiceHandler(w http.ResponseWriter, r *http.Request) {
	// Generate dice rolls and output messages
	val := rand.Intn(6) + 1
	str := fmt.Sprintf("The dice rolls are %d", val)
	w.Write([]byte(str))
}