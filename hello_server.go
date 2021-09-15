package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloHandler) // Register a handler
	http.ListenAndServe(":8888", nil) // Start the server
}

// HelloHandler  Describes what the server does.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
