package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Com _ na frente, significa que não é esse arquivo que vai usar o pacote, no nosso caso seria o "database/sql"
)

func main() {
	stringConexao := "root:gomes@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
