package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	//Quando passa o parametro normal passa como copia, quando com ponteiro passa uma referencia
	//Com ponteiro muda o valor da variavel
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
