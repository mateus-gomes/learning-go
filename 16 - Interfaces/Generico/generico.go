package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	//CUIDADO PARA NÃO FAZER GAMBIARRA!
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "String", false, true, float64(12345)) //Println é um exemplo de uso de interface generica
}
