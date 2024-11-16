package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Book" + id))
	})
	mux.HandleFunc("GET /books/dir/{d...}", func(w http.ResponseWriter, r *http.Request) {
		dirpath := r.PathValue("d")
		fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
	})
	mux.HandleFunc("GET /books/{$}", func(w http.ResponseWriter, r *http.Request) { // pegar algo exato
		w.Write([]byte("Books"))
	})
	mux.HandleFunc("GET /books/precedence/latest", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Books Precedence"))
	})
	mux.HandleFunc("GET /books/precedence/{x}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Books Precedence 2"))
	})
	// mux.HandleFunc("GET /books/{s}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Books Precedence"))
	// })
	// mux.HandleFunc("GET /{s}/latest", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Books Precedence 2"))
	// })
	http.ListenAndServe(":8080", mux)
}
