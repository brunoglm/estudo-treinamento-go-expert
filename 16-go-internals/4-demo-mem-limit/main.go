package main

import (
	"runtime"
	"time"
)

func main() {
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	for range 10 {
		allocateMemory(20 * 1024 * 1024) // Aloca 20MB de memória
		// Exibindo o uso de memória após as alocações
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		println("Alocação total de memória:", m.Alloc/1024/1024, "MB")
		println("Total de memória alocada:", m.TotalAlloc/1024/1024, "MB")
		println("Memória do sistema:", m.Sys/1024/1024, "MB")
		println("Número de buscas de ponteiros:", m.Lookups)
		println("Número de alocações de memória:", m.Mallocs)
		println("Número de liberações de memória:", m.Frees)
		println("Memória heap alocada:", m.HeapAlloc/1024/1024, "MB")
		println("Memória heap total:", m.HeapSys/1024/1024, "MB")
		println("Memória heap inativa:", m.HeapIdle/1024/1024, "MB")
		println("Memória heap em uso:", m.HeapInuse/1024/1024, "MB")
		println("Memória heap liberada:", m.HeapReleased/1024/1024, "MB")
		println("Número de objetos heap:", m.HeapObjects)
		println("Memória stack em uso:", m.StackInuse/1024/1024, "MB")
		println("Memória stack total:", m.StackSys/1024/1024, "MB")
		println("Memória MSpan em uso:", m.MSpanInuse/1024/1024, "MB")
		println("Memória MSpan total:", m.MSpanSys/1024/1024, "MB")
		println("Memória MCache em uso:", m.MCacheInuse/1024/1024, "MB")
		println("Memória MCache total:", m.MCacheSys/1024/1024, "MB")
		println("Memória BuckHash total:", m.BuckHashSys/1024/1024, "MB")
		println("Memória GC total:", m.GCSys/1024/1024, "MB")
		println("Memória OtherSys total:", m.OtherSys/1024/1024, "MB")
		println("Número de coletas de lixo:", m.NumGC)
		time.Sleep(1 * time.Second)
	}
}

// GODEBUG=gctrace=1 significa que o programa Go irá imprimir informações detalhadas sobre as operações de coleta de lixo (GC) que ocorrem durante a execução do programa. Isso inclui quando o GC é iniciado, quanto tempo leva para concluir, quantos objetos foram coletados e outras estatísticas relacionadas à memória.
// GOGC=300 significa que o coletor de lixo do Go será acionado quando a quantidade de memória alocada atingir 300% da quantidade de memória atualmente em uso. Isso pode resultar em uma coleta de lixo mais agressiva, o que pode ser útil para observar o comportamento do GC em um programa que aloca muita memória.
// GOGC=-1 desliga o garbage collector
// GOMEMLIMIT=5MiB limita a quantidade de memória soft, que o GC vai tentar manter abaixo desse limite. Se a alocação de memória ultrapassar esse limite, o GC será acionado para tentar liberar memória e reduzir o uso total. Isso pode ser útil para controlar o consumo de memória em programas que podem crescer rapidamente ou que precisam operar dentro de restrições de memória específicas.
// GODEBUG=gctrace=1 GOGC=80 GOMEMLIMIT=5MiB go run main.go
