package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	//... utilizado para demonstrar que o array tem o tamanho equivalente a quantidade de dados
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice não tem limite de valores
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//append cria um novo slice com mais um item
	slice = append(slice, 18)
	fmt.Println(slice)

	//Primeiro valor é incluso e ultimo valor é excluido
	//funciona como um ponteiro, então caso o array seja alterado o slice tbm sera
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays internos
	//Funciona como se fosse um Array com o tamanho da capacidade, porém você tem um slice com o tamanho que você passou, no exemplo seria 10
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	//Sempre que o array interno chegar no limite ele pega a capacidade e duplica
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
