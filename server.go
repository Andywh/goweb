package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Response) {
		w.Write([]byte("Hello, this is version 1!"))
	})
}
