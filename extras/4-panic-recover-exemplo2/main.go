package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		fmt.Println("Entrou no Handler")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic-handler")
	})

	addr := ":8080"
	log.Println("Listening on", addr)
	if err := http.ListenAndServe(addr, recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", addr, err)
	}
}
