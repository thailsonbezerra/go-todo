package router

import (
	"github.com/gorilla/mux"
	"github.com/thailsonbezerra/go-todo/src/controllers"
)

func Router(r *mux.Router) {
	r.HandleFunc("/todos", controllers.CriarTodo).Methods("POST")
	r.HandleFunc("/todos", controllers.BuscarTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.BuscarTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.AtualizarTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", controllers.DeletarTodo).Methods("DELETE")
}