package main

import (
	"errors"
	"fmt"
)

func main() {
	//Se só usar o INT ele pega como base a arquitetura do seu Computador

	//INICIO NUMEROS INTEIROS
	var numero int16 = 100
	fmt.Println(numero)

	//uint = sem sinais como "-"
	var numero2 uint32 = 10000
	fmt.Print(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//FIM NUMEROS INTEIROS

	//INICIO NUMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	//INICIO DAS STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := "A"
	fmt.Println(char)

	//FIM DAS STRINGS

	//TODA VARIAVEL TEM SEU VALOR BASE PARA STRING É " " PARA INT 0 PARA BOOLEANO FALSE E PARA TIPO ERROR É <nil>
	var texto int16
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
