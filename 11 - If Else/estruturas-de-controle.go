package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init
	//A variavel só é usada no escopo do if, depois ela morre
	if outronumero := numero; outronumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número entre 0 e -10")
	}
}
