package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var leakedGoroutines int32

func leakGoroutine() {
	atomic.AddInt32(&leakedGoroutines, 1)
	for {
		// loop infinito representando uma goroutine vazando recursos
	}
}

func leakHandler(w http.ResponseWriter, _ *http.Request) {
	go leakGoroutine() // Inicia uma goroutine que vaza recursos
	fmt.Fprintf(w, "Leaked goroutines!\n")
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	// Exibe o número atual de goroutines vazadas
	fmt.Fprintf(w, "Current number of leaked goroutines: %d\n", atomic.LoadInt32(&leakedGoroutines))
}

func main() {
	http.HandleFunc("/leak", leakHandler)     // Endpoint para vazar goroutines
	http.HandleFunc("/status", statusHandler) // Endpoint para verificar o status das goroutines vazadas

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// GOMAXPROCS=1 significa que o programa Go usará apenas um núcleo de CPU.
// GODEBUG=schedtrace=1 ativa o rastreamento do agendador, permitindo que você veja informações detalhadas sobre a execução das goroutines, incluindo quando elas são criadas, bloqueadas ou terminadas. Isso é útil para identificar vazamentos de goroutines e entender o comportamento do agendador em seu programa.
// GOMAXPROCS=1 GODEBUG=schedtrace=1 go run main.go

// SCHED 52076ms: gomaxprocs=1 idleprocs=1 threads=6 spinningthreads=0 needspinning=0 idlethreads=2 runqueue=0 [ 0 ] schedticks=[ 7 ]
// Explicando cada item do output:
// - SCHED: Indica que esta linha é uma saída do rastreamento do agendador.
// - 52076ms: O tempo decorrido desde o início do programa, em milissegundos.
// - gomaxprocs=1: O número de núcleos de CPU que o programa está usando (definido por GOMAXPROCS).
// - idleprocs=1: O número de processadores que estão ociosos (não executando goroutines).
// - threads=6: O número total de threads do sistema operacional que estão sendo usadas pelo programa Go.
// - spinningthreads=0: O número de threads que estão ativamente tentando encontrar trabalho para fazer (spinning).
// - needspinning=0: O número de threads que precisam começar a girar para encontrar trabalho, mas ainda não começaram.
// - idlethreads=2: O número de threads que estão ociosas e não estão fazendo nada.
// - runqueue=0: O número de goroutines que estão esperando para serem executadas (na fila de execução).
// - [ 0 ]: A lista de processadores e suas respectivas goroutines em execução (neste caso, o processador 0 está vazio).
// - schedticks=[ 7 ]: O número de ticks do agendador que ocorreram desde o início do programa.
