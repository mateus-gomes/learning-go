package main

import "fmt"

//Todas funções recursivas precisam falar onde termina
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(15)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
