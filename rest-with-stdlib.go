package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, rest with standard lib."))
}
