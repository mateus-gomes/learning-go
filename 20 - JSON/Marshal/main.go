package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint8  `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroEmJSON2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON2)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON2))
}
