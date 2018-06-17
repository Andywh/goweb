package main

import (
	"net/http"
	"log"
	"time"
	"os/signal"
	"os"
)


type myHandler struct{}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, the request URL is: " + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("Bye bye, this is version 3!--"))
}

func main() {
	server := &http.Server {
		Addr: ":4000",
		WriteTimeout: 4 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	server.Handler = mux

	go func() {
		<-quit

		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting server... v3")
	log.Fatal(server.ListenAndServe())
	err := server.ListenAndServe()
	if err != nil {
		if http.ErrServerClosed == err {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
	log.Println("Server exit!")
}
