package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//Quanto mais goroutines, mais rapido
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
