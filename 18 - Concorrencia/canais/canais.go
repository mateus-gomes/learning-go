package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //chan == o tipo do canal, ou seja ele só pode trafegar esse tipo de dados no canal
	go escrever("Olá Mundo!", canal)

	//Uma maneira de fazer o canal esperar
	// for {
	// 	mensagem, aberto := <-canal //esperando que o canal receba um valor
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //Mandando valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
