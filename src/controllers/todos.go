package controllers

import (
	"net/http"
)

func CriarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando um TODO!"))
}

func BuscarTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos os TODOs!"))
}

func BuscarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um TODO!"))
}

func AtualizarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando um TODO!"))
}

func DeletarTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um TODO!"))
}