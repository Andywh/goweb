package main

import (
	"net/http"
	"log"
)


type myHandler struct{}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this is version 2!"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye, this is version 2!"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting server... v2")
	log.Fatal(http.ListenAndServe(":4001", mux))
}
