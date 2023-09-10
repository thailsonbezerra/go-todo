package main

import (
	"fmt"
	"log"

	"github.com/thailsonbezerra/go-todo/database"
	"github.com/thailsonbezerra/go-todo/models/todos"
)

func main() {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conexão com o banco de dados bem-sucedida!")

	todo := todos.Todo{ID: 1, Created_at: "2023-09-08",Title: "Teste", Description: "Teste", Completed: false}
	fmt.Print(todo)

}