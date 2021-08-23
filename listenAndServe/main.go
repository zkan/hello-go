package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
