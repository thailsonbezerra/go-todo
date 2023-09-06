package main

import (
	"fmt"
	"log"

	"github.com/thailsonbezerra/go-todo/database"
)

func main() {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conexão com o banco de dados bem-sucedida!")
}