package main

import (
	"net/http"
	"log"
	"time"
)


type myHandler struct{}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, the request URL is: " + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye, this is version 3!"))
}

func main() {
	server := &http.Server {
		Addr:":4000",
		WriteTimeout:2*time.Second,
	}

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	server.Handler = mux

	log.Println("Starting server... v3")
	//log.Fatal(http.ListenAndServe(":4000", mux))
	log.Fatal(server.ListenAndServe())
}
