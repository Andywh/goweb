package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("宋亮，这是一个测试的程序!"))
	})

	log.Println("Starting server... v1")
	log.Fatal(http.ListenAndServe(":80", nil))
}
