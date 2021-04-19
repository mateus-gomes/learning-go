package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//Go não deixa fazer contas com tipos diferentes nem int16 e int32
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Operadores logicos
	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero -= 15
	fmt.Println(numero)
	numero *= 15
	fmt.Println(numero)
	numero /= 15
	fmt.Println(numero)
	numero %= 15
	fmt.Println(numero)
}
