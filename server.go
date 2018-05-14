package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is version 1!"))
	})

	log.Println("Starting server... v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
