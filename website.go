package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there Im Rizki")
}

func main() {
	http.HandleFunc("/" , handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}