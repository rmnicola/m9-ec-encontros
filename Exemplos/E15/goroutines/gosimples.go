package main

import (
	"fmt"
	"time"
)

// Função que exibe uma mensagem após um sleep de 2 segundos
func printWithDelay(msg string) {
	time.Sleep(2 * time.Second)
	fmt.Println(msg)
}

func main() {
	start := time.Now() // Início da medição do tempo

	// Chamando a função 3 vezes de forma concorrente
	go printWithDelay("Mensagem 1")
	go printWithDelay("Mensagem 2")
	go printWithDelay("Mensagem 3")

	elapsed := time.Since(start) // Tempo total de execução
	fmt.Printf("Tempo total de execução: %s\n", elapsed)
}
