package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //Avisa que tem 2 funções para ser executadas

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 função a ser executada, resta 1 função
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 função a ser executada, todas foram executadas
	}()

	waitGroup.Wait() //Pede para esperar
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
