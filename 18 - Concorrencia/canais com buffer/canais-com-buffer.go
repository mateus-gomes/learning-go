package main

import "fmt"

func main() {
	canal := make(chan string, 2) //Só bloqueia se tiver chegado no máximo
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
