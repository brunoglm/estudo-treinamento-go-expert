package main

import (
	"fmt"
	"net/http"
)

var number uint64 = 0

func main() {
	// mu := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// mu.Lock()
		number++
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes\n", number)))
		// mu.Unlock()
	})
	http.ListenAndServe(":8080", nil)
}
