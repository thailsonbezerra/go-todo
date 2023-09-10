package main

import (
	"fmt"
	"log"

	"github.com/thailsonbezerra/go-todo/src/config/database"
	"github.com/thailsonbezerra/go-todo/src/models/todos"
)

func main() {
	db, err := database.OpenConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todo := todos.Todo{ID: 1, CreatedAt: "2023-09-08",Title: "Teste", Description: "Teste", Completed: false}
	fmt.Print(todo)

}