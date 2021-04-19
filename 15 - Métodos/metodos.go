package main

import "fmt"

//A DIFERENÇA ENTRE MÉTODO E FUNÇÃO
//Basicamente a diferença esta no método estar atrelado à uma struct

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 31}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
