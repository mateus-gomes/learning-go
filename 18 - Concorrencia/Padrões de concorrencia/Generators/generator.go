package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")
	//Utilizado para abstrair a complexidade da quantidade de chamadas de goroutines

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
