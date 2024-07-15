package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// mu := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// mu.Lock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes\n", number)))
		// mu.Unlock()
	})
	http.ListenAndServe(":8080", nil)
}
