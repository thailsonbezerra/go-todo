package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/thailsonbezerra/go-todo/src/config/database"
	"github.com/thailsonbezerra/go-todo/src/models/todos"
	"github.com/thailsonbezerra/go-todo/src/router"
)

func main() {
	r := mux.NewRouter()
	router.Router(r)


	db, err := database.OpenConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	todo := todos.Todo{ID: 1, CreatedAt: "2023-09-08",Title: "Teste", Description: "Teste", Completed: false}
	fmt.Print(todo)

	http.ListenAndServe(":3000", r)

}