package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	go escrever("Olá Mundo!") //goroutine = Independente de ter terminado essa função segue o programa
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
