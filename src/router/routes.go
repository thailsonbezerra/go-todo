package router

import (
	"github.com/gorilla/mux"
	"github.com/thailsonbezerra/go-todo/src/controllers/todos"
)

func Router(r *mux.Router) {
	r.HandleFunc("/todos", todos.CriarTodo).Methods("POST")
	r.HandleFunc("/todos", todos.BuscarTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", todos.BuscarTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todos.AtualizarTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todos.DeletarTodo).Methods("DELETE")
}
